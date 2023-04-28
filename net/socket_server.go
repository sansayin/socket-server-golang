package net

import (
	"log"
	"net"
	"pattern/utils"
	"sync"
	"time"
)

const SOCKET_TIMEOUT = 2

var (
	L    = log.Printf //fmt.Printf
	void = struct{}{}
)

type IServant interface {
	OnRequest(net.Conn, []byte) []byte
}

type SocketServer struct {
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

func NewSocketServer(max int, long bool, debug bool) *SocketServer {
	maxChan := make(chan struct{}, max)
	for i := 0; i < max; i++ {
		maxChan <- void
	}
	servants := make(map[IServant]struct{}, 0)
	clients := utils.NewClientDict()
	return &SocketServer{
		max_rountins:    maxChan,
		servants:        servants,
		long_connection: long,
		clients:         clients,
		debug:           debug,
	}
}

func (ss *SocketServer) NumGoRoutine() int {
	return len(ss.clients.Get())
}

func (ss *SocketServer) StartTCP(ipAddr string, port string) error {
	ss.ip_addr, ss.port = ipAddr, port
	server, err := net.Listen("tcp", ss.ip_addr+":"+ss.port)
	if err != nil {
		panic(err.Error())
	}
	defer server.Close()
	L("TCP Server Started:" + ss.ip_addr + " on port:" + ss.port)
	ss.tcp_server = server
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				ss.max_rountins <- void
			}
		}
	}()
	for {
		connection, err := server.Accept()
		if err != nil {
			L("Error accepting: %v\n", err.Error())
			continue
		}
		//L("Client: %v connected", connection.RemoteAddr())
		<-ss.max_rountins
		go ss.processTcpClient(connection, done)
	}
}

func (ss *SocketServer) processTcpClient(conn net.Conn, done chan bool) {
	ss.clients.Add(&conn)
	buffer := make([]byte, 1024)

	defer func() {
		//L("Client: %v left", conn)
		ss.clients.Del(&conn)
		done <- true
	}()
	//use loop for long-connection-socket case
	for {
		if ss.long_connection && !ss.debug { //still need timeout when client silence too long
			conn.SetReadDeadline(time.Now().Add(SOCKET_TIMEOUT * time.Second))
		}

		length, err := conn.Read(buffer)
		if err != nil {
			//L("Read Error: %s : buffer-length: %d\n", err, length)
			break
		}
		for k := range ss.servants {
			reply := k.OnRequest(conn, buffer[0:length])
			conn.Write([]byte(reply))
		}
		if ss.stop || !ss.long_connection {
			break
		}
	}
}

func (ss *SocketServer) StartUDP(port string) error {
	ss.port = port
	server, err := net.ListenPacket("udp", ":"+ss.port)
	if err != nil {
		return err
	}
	defer server.Close()
	L("UDP Server Started: on port:" + ss.port)
	ss.udp_server = server
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				ss.max_rountins <- void
			}
		}
	}()
	for {
		buf := make([]byte, 1024)
		length, addr, err := server.ReadFrom(buf)
		if err != nil {
			continue
		}
		<-ss.max_rountins
		go ss.processUdpClient(addr, buf[0:length], done)
	}
}

func (ss *SocketServer) processUdpClient(addr net.Addr, buf []byte, done chan bool) {
	defer func() {
		done <- true
	}()
	for k := range ss.servants {
		reply := k.OnRequest(nil, buf)
		ss.udp_server.WriteTo(reply, addr)
	}
}

func (ss *SocketServer) AddServant(servant interface{ IServant }) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	ss.servants[servant] = void
}

func (ss *SocketServer) RemoveServant(servant interface{ IServant }) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	delete(ss.servants, servant)
}

func (ss *SocketServer) BroadCast(msg string) {
	for c, v := range ss.clients.Get() {
		if v {
			_, err := (*c).Write([]byte(msg))
			if err != nil {
				L("Broadcast Error: %v", err)
			}
		}
	}
}

func (ss *SocketServer) Stop() <-chan time.Time {
	//		ss.stop = true
	L("Shuting down ...")
	for servant := range ss.servants {
		ss.RemoveServant(servant)
	}
	ss.BroadCast("\n**************Going Offline**************\r\n\r\n")
	L("Closing %d connections", len(ss.clients.Get()))
	for c, v := range ss.clients.Get() {
		if v {
			(*c).Close()
		}
		ss.clients.Del(c)
	}
	return time.After(1 * time.Second)
}
