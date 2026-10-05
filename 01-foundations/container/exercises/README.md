# container/list, container/heap, container/ring exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/container/exercises/...
```

## Exercise: LRU cache

Implement a fixed-capacity least-recently-used cache:

```go
func NewLRU(capacity int) *LRU
func (c *LRU) Get(key string) (int, bool)
func (c *LRU) Put(key string, val int)
func (c *LRU) Len() int
```

Rules:

- `Get` marks the entry as most recently used
- `Put` on an existing key updates the value and marks it recently used
- When full, `Put` of a new key evicts the least recently used entry
- Capacity `0` stores nothing

Hints: a `map[string]*list.Element` for O(1) lookup plus a `list.List` for
usage order, with `MoveToFront` on every access.

## Bonus
Make it generic: `LRU[K comparable, V any]`.
