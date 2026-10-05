// Package solution holds the reference solutions for the sync/atomic exercises.
package solution

import "sync/atomic"

// Stats counts requests and failures without locks.
type Stats struct {
	requests atomic.Int64
	failures atomic.Int64
}

// Record counts one request, and one failure when failed is true.
func (s *Stats) Record(failed bool) {
	s.requests.Add(1)
	if failed {
		s.failures.Add(1)
	}
}

// Snapshot returns the current counts.
func (s *Stats) Snapshot() (requests, failures int64) {
	return s.requests.Load(), s.failures.Load()
}

// Config is an immutable application setting.
type Config struct {
	Name  string
	Limit int
}

// ConfigStore lets many goroutines read the config while others replace it.
type ConfigStore struct {
	p atomic.Pointer[Config]
}

// NewConfigStore creates a store holding c.
func NewConfigStore(c Config) *ConfigStore {
	s := &ConfigStore{}
	s.p.Store(&c)
	return s
}

// Get returns the current config.
func (s *ConfigStore) Get() Config { return *s.p.Load() }

// Update applies f to the current config atomically. f may be called more than
// once if another update races with it, so it must have no side effects.
func (s *ConfigStore) Update(f func(Config) Config) {
	for {
		old := s.p.Load()
		updated := f(*old)
		if s.p.CompareAndSwap(old, &updated) {
			return
		}
	}
}
