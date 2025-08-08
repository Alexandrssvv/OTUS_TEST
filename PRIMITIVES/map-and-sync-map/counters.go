package main

import "sync"

type Counters struct {
	mutex sync.Mutex
	data  map[int]int
}

func NewCounters() *Counters {
	return &Counters{
		sync.Mutex{},
		make(map[int]int),
	}
}

func (c *Counters) Load(key int) (int, bool) {
	c.mutex.Lock()

	val, ok := c.data[key]
	return val, ok
}

func (c *Counters) Store(key int, value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
}

func (c *Counters) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data = make(map[int]int)
}
