package utils

import (
	"net"
	"sync"
)

type ClientDict struct {
	mutex sync.RWMutex
	dict  map[*net.Conn]bool
}

func NewClientDict() *ClientDict {
	return &ClientDict{
		dict: make(map[*net.Conn]bool, 0),
	}
}

func (sc *ClientDict) Get() map[*net.Conn] bool{
  sc.mutex.RLock()
  defer sc.mutex.RUnlock()
  return sc.dict
}

func (sc *ClientDict) Add(c *net.Conn) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	sc.dict[c] = true
}

func (sc *ClientDict) Del(c *net.Conn) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	delete(sc.dict, c)
}
