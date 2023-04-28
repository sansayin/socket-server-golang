package net

import (
	"net"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpServant_OnRequest(t *testing.T) {
	type fields struct {
		Id         int
		StaticRoot string
		bufferPool sync.Pool
		mime       sync.Map
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
			h := &HttpServant{
				Id:         tt.fields.Id,
				StaticRoot: tt.fields.StaticRoot,
				bufferPool: tt.fields.bufferPool,
				mime:       tt.fields.mime,
			}
			if got := h.OnRequest(tt.args.conn, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpServant.OnRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpServant_getRequestingFile(t *testing.T) {
	type fields struct {
		Id         int
		StaticRoot string
		bufferPool sync.Pool
		mime       sync.Map
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HttpServant{
				Id:         tt.fields.Id,
				StaticRoot: tt.fields.StaticRoot,
				bufferPool: tt.fields.bufferPool,
				mime:       tt.fields.mime,
			}
			if got,_ := h.getRequestingFile(tt.args.buf); got != tt.want {
				t.Errorf("HttpServant.getRequestingFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpServant_getFileContentType(t *testing.T) {
	type fields struct {
		Id         int
		StaticRoot string
		bufferPool sync.Pool
		mime       sync.Map
	}
	type args struct {
		request_file string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
      name: "HTML File", 
      fields: fields{Id: 1, StaticRoot: "../static/"}, 
      args: args{request_file: "index.html"}, 
      want: "text/html", 
      wantErr: false,
    },
		{
      name: "Image File", 
      fields: fields{Id: 1, StaticRoot: "../static/"}, 
      args: args{request_file: "wordcloud.png"}, 
      want: "image/png", 
      wantErr: false,
    },
		{
      name: "JSON File", 
      fields: fields{Id: 1, StaticRoot: "../static/"}, 
      args: args{request_file: "package.json"}, 
      want: "application/json", 
      wantErr: false,
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HttpServant{
				Id:         tt.fields.Id,
				StaticRoot: tt.fields.StaticRoot,
			}
			got, _ := h.getFileContentType(tt.args.request_file)
			assert.Equal(t, got, tt.want)
		})
	}
}
