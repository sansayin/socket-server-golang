package utils

import (
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{v: make(map[string]int)}
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	s.v[key]++
	s.mu.Unlock()
}

func (s *SafeCounter) Get(key string) int {
	return s.v[key]
}
