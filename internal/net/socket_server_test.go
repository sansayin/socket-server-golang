package net

import (
	"net"
	"reflect"
	"testing"
)

func Test_defaultOpts(t *testing.T) {
	tests := []struct {
		name string
		want Opts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultOpts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMaxRoutines(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want OptsFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMaxRoutines(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxRoutines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLongConn(t *testing.T) {
	type args struct {
		long bool
	}
	tests := []struct {
		name string
		args args
		want OptsFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLongConn(tt.args.long); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLongConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDebug(t *testing.T) {
	type args struct {
		debug bool
	}
	tests := []struct {
		name string
		args args
		want OptsFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDebug(tt.args.debug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPrometheus(t *testing.T) {
	tests := []struct {
		name string
		want OptsFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPrometheus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPrometheus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSocketServer(t *testing.T) {
	type args struct {
		opts []OptsFunc
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
			if got := NewSocketServer(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSocketServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocketServer_NumGoRoutine(t *testing.T) {
	type fields struct {
		Opts Opts
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
				Opts: tt.fields.Opts,
			}
			if got := ss.NumGoRoutine(); got != tt.want {
				t.Errorf("SocketServer.NumGoRoutine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocketServer_StartTCP(t *testing.T) {
	type fields struct {
		Opts Opts
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
				Opts: tt.fields.Opts,
			}
			if err := ss.StartTCP(tt.args.ipAddr, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("SocketServer.StartTCP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSocketServer_processTcpClient(t *testing.T) {
	type fields struct {
		Opts Opts
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
				Opts: tt.fields.Opts,
			}
			ss.processTcpClient(tt.args.conn, tt.args.done)
		})
	}
}

func TestSocketServer_StartUDP(t *testing.T) {
	type fields struct {
		Opts Opts
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
				Opts: tt.fields.Opts,
			}
			if err := ss.StartUDP(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("SocketServer.StartUDP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSocketServer_processUdpClient(t *testing.T) {
	type fields struct {
		Opts Opts
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
				Opts: tt.fields.Opts,
			}
			ss.processUdpClient(tt.args.addr, tt.args.buf, tt.args.done)
		})
	}
}

func TestSocketServer_AddServant(t *testing.T) {
	type fields struct {
		Opts Opts
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
				Opts: tt.fields.Opts,
			}
			ss.AddServant(tt.args.servant)
		})
	}
}

func TestSocketServer_RemoveServant(t *testing.T) {
	type fields struct {
		Opts Opts
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
				Opts: tt.fields.Opts,
			}
			ss.RemoveServant(tt.args.servant)
		})
	}
}

func TestSocketServer_BroadCast(t *testing.T) {
	type fields struct {
		Opts Opts
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
				Opts: tt.fields.Opts,
			}
			ss.BroadCast(tt.args.msg)
		})
	}
}

func TestSocketServer_Stop(t *testing.T) {
	type fields struct {
		Opts Opts
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SocketServer{
				Opts: tt.fields.Opts,
			}
			ss.Stop()
		})
	}
}
