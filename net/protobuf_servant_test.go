package net

import (
	"net"
	"reflect"
	"testing"
)

func TestProtoBufServant_OnRequest(t *testing.T) {
	type fields struct {
		Id int
	}
	type args struct {
		conn net.Conn
		msg  []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProtoBufServant{
				Id: tt.fields.Id,
			}
			if got := p.OnRequest(tt.args.conn, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProtoBufServant.OnRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
