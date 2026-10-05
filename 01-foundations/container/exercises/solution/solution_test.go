package solution

import "testing"

func TestLRU(t *testing.T) {
	c := NewLRU(2)
	c.Put("a", 1)
	c.Put("b", 2)
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Fatalf("Get(a) = %d, %v", v, ok)
	}
	c.Put("c", 3) // evicts b, the least recently used

	if _, ok := c.Get("b"); ok {
		t.Error("b should have been evicted")
	}
	if _, ok := c.Get("a"); !ok {
		t.Error("a should still be cached")
	}
	if c.Len() != 2 {
		t.Errorf("Len() = %d, want 2", c.Len())
	}

	c.Put("a", 10) // update, not insert
	if v, _ := c.Get("a"); v != 10 || c.Len() != 2 {
		t.Errorf("after update: v=%d len=%d", v, c.Len())
	}
}

func TestLRUZeroCapacity(t *testing.T) {
	c := NewLRU(0)
	c.Put("a", 1)
	if c.Len() != 0 {
		t.Errorf("Len() = %d, want 0", c.Len())
	}
}
