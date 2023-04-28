package utils

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewSafeCounter(t *testing.T) {
	tests := []struct {
		name string
		want *SafeCounter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSafeCounter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSafeCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeCounter_Inc(t *testing.T) {
	type fields struct {
		mu sync.Mutex
		v  map[string]int
	}
	type args struct {
		key string
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
			s := &SafeCounter{
				mu: tt.fields.mu,
				v:  tt.fields.v,
			}
			s.Inc(tt.args.key)
		})
	}
}

func TestSafeCounter_Get(t *testing.T) {
	type fields struct {
		mu sync.Mutex
		v  map[string]int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SafeCounter{
				mu: tt.fields.mu,
				v:  tt.fields.v,
			}
			if got := s.Get(tt.args.key); got != tt.want {
				t.Errorf("SafeCounter.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
