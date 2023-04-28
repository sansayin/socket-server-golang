package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"pattern/net"
	"runtime"
	"runtime/trace"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

const (
	SERVER_HOST       = "0.0.0.0"
	SERVER_PORT       = "9988"
	MAX_PROC_ROUTINES = 2000
	SOCKET_TIMEOUT    = 100
	LONG_CONNECTION   = true 
	DEBUG_MODE        = true
)

var L = log.Printf

var server *net.SocketServer

// Implement Interface from net.SocketServer
func handler(signal os.Signal) {
	if signal == syscall.SIGTERM || signal == syscall.SIGINT {
		if server != nil {
			<-server.Stop()
		}
		os.Exit(0)
	} /* else {
		L("Ignoring signal: %v", signal)
	}*/
}

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
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)


  /******Metric Start*********/
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	socket_metrics := NewMetrics(reg)
  socket_metrics.devices.Set(0)

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	http.Handle("/metrics", promHandler)
	go http.ListenAndServe(":8081", nil)

  /******Metric End*********/

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case s := <-sigchnl:
				handler(s)
			case <-ticker.C:
				L("Running Routins: %v\n", runtime.NumGoroutine())
			}
		}
	}()

	server = net.NewSocketServer(
		MAX_PROC_ROUTINES,
		LONG_CONNECTION,
		DEBUG_MODE,
	)

	//server.AddServant(&net.ProtoBufServant{Id: 1})
	server.AddServant(&net.HttpServant{Id: 2, StaticRoot: "./static"})

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
