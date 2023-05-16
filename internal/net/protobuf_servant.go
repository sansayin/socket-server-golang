package net

import (
	"fmt"
	"net"
	"pattern/message/pb"

	"github.com/golang/protobuf/proto"
)

type ProtoBufServant struct {
	Id int
}

func (p ProtoBufServant) OnRequest(conn net.Conn, msg []byte) []byte {
	var msgPb pb.Message

	if err := proto.Unmarshal(msg, &msgPb); err != nil {
		return nil
	}
	responseStr := fmt.Sprintf("Received message: #%v", msgPb.GetId())
	return []byte(responseStr)
}
