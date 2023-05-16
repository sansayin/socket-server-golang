package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"pattern/internal/net"
	"runtime"
	"runtime/trace"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	SERVER_HOST       = "0.0.0.0"
	SERVER_PORT       = "9988"
	MAX_PROC_ROUTINES = 20000
	SOCKET_TIMEOUT    = 100
	LONG_CONNECTION   = true
	DEBUG_MODE        = true
)

var L = log.Printf

var server *net.SocketServer

var (
	bProf  bool
	bTrace bool
)

type metrics struct {
	devices prometheus.Gauge
	info    *prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "socket_server",
			Name:      "connected_devices",
			Help:      "Number of currently connected devices.",
		}),
	}
	reg.MustRegister(m.devices)
	return m
}

func initMonitor() {

}

func main() {
	flag.BoolVar(&bProf, "pprof", false, "turn on pprof")
	flag.BoolVar(&bTrace, "trace", false, "trace go routines")
	flag.Parse()

	if bProf {
		go func() {
			log.Println(http.ListenAndServe(":6060", nil))
		}()
	}
	if bTrace {
		trace.Start(os.Stderr)
		defer trace.Stop()
	}

	/******Metric Start*********/
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	socket_metrics := NewMetrics(reg)
	socket_metrics.devices.Set(0)

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	http.Handle("/metrics", promHandler)
	go http.ListenAndServe(":8081", nil)

	/******Metric End*********/

	server := net.NewSocketServer(net.WithMaxRoutines(MAX_PROC_ROUTINES), net.WithDebug(true), net.WithLongConn(true))
	//server.AddServant(&net.ProtoBufServant{Id: 1})
	server.AddServant(&net.HttpServant{Id: 2, StaticRoot: "./static"})

	ticker := time.NewTicker(10 * time.Second)
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				server.Stop()
				stop()
                os.Exit(0)
			case <-ticker.C:
				L("Running Routins: %v\n", runtime.NumGoroutine())
			}
		}
	}()

	if true {
		if err := server.StartTCP(SERVER_HOST, SERVER_PORT); err != nil {
			panic(err)
		}
	} else {
		if err := server.StartUDP(SERVER_PORT); err != nil {
			panic(err)
		}
	}
}