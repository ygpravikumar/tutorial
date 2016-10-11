package safecounter

import "sync"

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

//CreateCounter to create and return a new safecounter objct
func CreateCounter() SafeCounter {
	return SafeCounter{v: make(map[string]int)}
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	val, ok := c.v[key]
	if ok == true {
		c.v[key] = val + 1
	} else {
		c.v[key] = 1
	}
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	val, ok := c.v[key]
	if ok == true {
		return val
	}

	return 0
}
