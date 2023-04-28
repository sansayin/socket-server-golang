package main

import (
	"net"
	"os"
	"pattern/message/pb"

	"github.com/golang/protobuf/proto"
)

func T_main() {
	udpServer, err := net.ResolveUDPAddr("udp", ":9988")

	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Listen failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()

	msg := pb.Message{
		Id:      1,
		Content: "hello",
	}
	data, err := proto.Marshal(&msg)
	if err != nil {
		println("marshaling error")
	}
	_, err = conn.Write(data)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println(string(received))
} 
