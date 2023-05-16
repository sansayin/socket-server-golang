package net

import (
	"net"
	"pattern/utils"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNewSocketServer(t *testing.T) {
	type args struct {
		max   int
		long  bool
		debug bool
	}
	tests := []struct {
		name string
		args args
		want *SocketServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSocketServer(tt.args.max, tt.args.long, tt.args.debug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSocketServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocketServer_NumGoRoutine(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			if got := ss.NumGoRoutine(); got != tt.want {
				t.Errorf("SocketServer.NumGoRoutine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocketServer_StartTCP(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	type args struct {
		ipAddr string
		port   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			if err := ss.StartTCP(tt.args.ipAddr, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("SocketServer.StartTCP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSocketServer_processTcpClient(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	type args struct {
		conn net.Conn
		done chan bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			ss.processTcpClient(tt.args.conn, tt.args.done)
		})
	}
}

func TestSocketServer_StartUDP(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	type args struct {
		port string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			if err := ss.StartUDP(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("SocketServer.StartUDP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSocketServer_processUdpClient(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	type args struct {
		addr net.Addr
		buf  []byte
		done chan bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			ss.processUdpClient(tt.args.addr, tt.args.buf, tt.args.done)
		})
	}
}

func TestSocketServer_AddServant(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	type args struct {
		servant interface{ IServant }
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			ss.AddServant(tt.args.servant)
		})
	}
}

func TestSocketServer_RemoveServant(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	type args struct {
		servant interface{ IServant }
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			ss.RemoveServant(tt.args.servant)
		})
	}
}

func TestSocketServer_BroadCast(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			ss.BroadCast(tt.args.msg)
		})
	}
}

func TestSocketServer_Stop(t *testing.T) {
	type fields struct {
		max_rountins    chan struct{}
		stop            bool
		ip_addr         string
		port            string
		tcp_server      net.Listener
		udp_server      net.PacketConn
		servants        map[IServant]struct{}
		long_connection bool
		clients         *utils.ClientDisct
		debug           bool
		mutex           sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   <-chan time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				max_rountins:    tt.fields.max_rountins,
				stop:            tt.fields.stop,
				ip_addr:         tt.fields.ip_addr,
				port:            tt.fields.port,
				tcp_server:      tt.fields.tcp_server,
				udp_server:      tt.fields.udp_server,
				servants:        tt.fields.servants,
				long_connection: tt.fields.long_connection,
				clients:         tt.fields.clients,
				debug:           tt.fields.debug,
				mutex:           tt.fields.mutex,
			}
			if got := ss.Stop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SocketServer.Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}
