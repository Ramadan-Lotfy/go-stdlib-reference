package solution

import (
	"sync"
	"testing"
)

func TestStats(t *testing.T) {
	var s Stats
	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Record(i%4 == 0)
		}()
	}
	wg.Wait()
	req, fail := s.Snapshot()
	if req != 100 || fail != 25 {
		t.Errorf("Snapshot() = %d, %d; want 100, 25", req, fail)
	}
}

func TestConfigStore(t *testing.T) {
	s := NewConfigStore(Config{Name: "app", Limit: 0})
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Update(func(c Config) Config { c.Limit++; return c })
		}()
	}
	wg.Wait()
	if got := s.Get(); got.Limit != 100 || got.Name != "app" {
		t.Errorf("Get() = %+v, want {app 100}", got)
	}
}
