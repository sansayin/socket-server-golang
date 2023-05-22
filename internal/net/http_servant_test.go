package net

import (
	"net"
	"reflect"
	"sync"
	"testing"
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
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
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
			got, err := h.getRequestingFile(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpServant.getRequestingFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
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
			got, err := h.getFileContentType(tt.args.request_file)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpServant.getFileContentType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HttpServant.getFileContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}
