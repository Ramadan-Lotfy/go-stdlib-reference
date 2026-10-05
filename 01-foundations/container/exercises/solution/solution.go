// Package solution holds the reference solutions for the container/list, container/heap, container/ring exercises.
package solution

import "container/list"

// LRU is a fixed-capacity least-recently-used cache.
type LRU struct {
	capacity int
	order    *list.List // front = most recently used
	items    map[string]*list.Element
}

type entry struct {
	key string
	val int
}

// NewLRU creates a cache holding at most capacity entries.
func NewLRU(capacity int) *LRU {
	return &LRU{capacity: capacity, order: list.New(), items: make(map[string]*list.Element)}
}

// Get returns the value for key and marks it as recently used.
func (c *LRU) Get(key string) (int, bool) {
	el, ok := c.items[key]
	if !ok {
		return 0, false
	}
	c.order.MoveToFront(el)
	return el.Value.(*entry).val, true
}

// Put stores a value, evicting the least recently used entry when full.
func (c *LRU) Put(key string, val int) {
	if c.capacity <= 0 {
		return
	}
	if el, ok := c.items[key]; ok {
		el.Value.(*entry).val = val
		c.order.MoveToFront(el)
		return
	}
	if c.order.Len() >= c.capacity {
		oldest := c.order.Back()
		c.order.Remove(oldest)
		delete(c.items, oldest.Value.(*entry).key)
	}
	c.items[key] = c.order.PushFront(&entry{key, val})
}

// Len returns the number of cached entries.
func (c *LRU) Len() int { return c.order.Len() }
