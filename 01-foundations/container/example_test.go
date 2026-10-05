package containerref_test

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

// Example_list builds a list from both ends and walks it.
func Example_list() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	var vals []any
	for e := l.Front(); e != nil; e = e.Next() {
		vals = append(vals, e.Value)
	}
	fmt.Println(vals, l.Len())

	l.Remove(e4)
	l.MoveToFront(l.Back()) // 3 moves to the front
	fmt.Println(l.Front().Value, l.Len())
	// Output:
	// [1 2 3 4] 4
	// 3 3
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Example_heap is a min-heap of ints. Pop always returns the smallest.
func Example_heap() {
	h := &IntHeap{5, 2, 8}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println("min:", (*h)[0])
	for h.Len() > 0 {
		fmt.Print(heap.Pop(h), " ")
	}
	fmt.Println()
	// Output:
	// min: 2
	// 2 3 5 8
}

type task struct {
	name     string
	priority int
}

type taskQueue []*task

func (q taskQueue) Len() int           { return len(q) }
func (q taskQueue) Less(i, j int) bool { return q[i].priority > q[j].priority } // max-heap
func (q taskQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *taskQueue) Push(x any)        { *q = append(*q, x.(*task)) }
func (q *taskQueue) Pop() any {
	old := *q
	n := len(old)
	t := old[n-1]
	*q = old[:n-1]
	return t
}

// Example_priorityQueue runs the highest-priority task first and re-prioritizes with Fix.
func Example_priorityQueue() {
	q := &taskQueue{}
	low := &task{"low", 1}
	for _, t := range []*task{{"mid", 5}, low, {"high", 9}} {
		heap.Push(q, t)
	}

	low.priority = 100 // changed in place, so restore the heap order
	for i, t := range *q {
		if t == low {
			heap.Fix(q, i)
		}
	}

	for q.Len() > 0 {
		fmt.Println(heap.Pop(q).(*task).name)
	}
	// Output:
	// low
	// high
	// mid
}

// Example_ring walks a circular list.
func Example_ring() {
	r := ring.New(3)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func(v any) { fmt.Println(v) })
	// Output:
	// 0
	// 1
	// 2
}
