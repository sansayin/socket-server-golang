package utils

import (
	"net"
	"sync"
)

type ClientDisct struct {
	mutex sync.RWMutex
	dict  map[*net.Conn]bool
}

func NewClientDict() *ClientDisct {
	return &ClientDisct{
		dict: make(map[*net.Conn]bool, 0),
	}
}

func (sc *ClientDisct) Get() map[*net.Conn] bool{
  sc.mutex.RLock()
  defer sc.mutex.RUnlock()
  return sc.dict
}

func (sc *ClientDisct) Add(c *net.Conn) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	sc.dict[c] = true
}

func (sc *ClientDisct) Del(c *net.Conn) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
  (*c).Close()
	delete(sc.dict, c)
}
