# container/list, container/heap, container/ring

## Overview
Three small data-structure packages. `list` is a doubly linked list,
`heap` turns any slice type into a priority queue, and `ring` is a circular
list. In everyday Go you mostly use slices and maps; reach for these when you
need O(1) removal from the middle (`list`, for example in an LRU cache) or
priority ordering (`heap`).

## Core API

| Function / Type | Purpose |
|---|---|
| `list.New`, `PushBack`, `PushFront` | Doubly linked list; `Element.Value` is `any` |
| `list.Remove`, `MoveToFront`, `InsertBefore` | O(1) edits given an `*Element` |
| `heap.Init`, `heap.Push`, `heap.Pop` | Operate on a type implementing `heap.Interface` |
| `heap.Fix`, `heap.Remove` | Restore order after changing or removing an element |
| `ring.New`, `Next`, `Do` | Circular list |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/container/
```

## Gotchas
- **`heap` is not a ready-made type.** You implement `Len`, `Less`, `Swap`, `Push` and `Pop` on your own slice type, then call the `heap.*` functions.
- **Call `heap.Push(h, x)` and `heap.Pop(h)`, not your own `h.Push` / `h.Pop`.** The methods you write only move data; the package restores heap order.
- **`Pop` on a `heap.Interface` removes the last slice element**; `heap.Pop` first swaps the root to the end.
- **`list` values are `any`**, so you need type assertions and lose type safety.
- **After changing a priority, call `heap.Fix`.**
- **A zero `list.List` is ready to use**, but never copy a list.

## Java comparison

| Java | Go |
|---|---|
| `PriorityQueue<T>` | `container/heap` on a slice type |
| `LinkedList<T>` | `container/list` |
| `LinkedHashMap` as LRU | `list` + `map` (see the exercise) |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/container/list, container/heap, container/ring
- Source: https://cs.opensource.google/go/go/+/master:src/container/list, container/heap, container/ring/
