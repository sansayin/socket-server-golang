package utils

import (
	"net"
	"reflect"
	"sync"
	"testing"
)

func TestNewClientDict(t *testing.T) {
	tests := []struct {
		name string
		want *ClientDisct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientDict(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientDisct_Get(t *testing.T) {
	type fields struct {
		mutex sync.RWMutex
		dict  map[*net.Conn]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[*net.Conn]bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &ClientDisct{
				mutex: tt.fields.mutex,
				dict:  tt.fields.dict,
			}
			if got := sc.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientDisct.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientDisct_Add(t *testing.T) {
	type fields struct {
		mutex sync.RWMutex
		dict  map[*net.Conn]bool
	}
	type args struct {
		c *net.Conn
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
			sc := &ClientDisct{
				mutex: tt.fields.mutex,
				dict:  tt.fields.dict,
			}
			sc.Add(tt.args.c)
		})
	}
}

func TestClientDisct_Del(t *testing.T) {
	type fields struct {
		mutex sync.RWMutex
		dict  map[*net.Conn]bool
	}
	type args struct {
		c *net.Conn
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
			sc := &ClientDisct{
				mutex: tt.fields.mutex,
				dict:  tt.fields.dict,
			}
			sc.Del(tt.args.c)
		})
	}
}
