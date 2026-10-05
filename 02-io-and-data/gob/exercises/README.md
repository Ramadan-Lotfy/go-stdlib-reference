# encoding/gob exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/gob/exercises/...
```

## Exercise: DeepCopy

Write a generic `DeepCopy[T any](src T) (T, error)` that returns a deep copy of
any gob-encodable value by encoding it to a buffer and decoding it into a new
value.

Check that changing the copy's slice and map does not change the original.

Hints: `bytes.Buffer`, `gob.NewEncoder`, `gob.NewDecoder`, and decode into a
fresh zero value of `T`.

## Bonus
Why would `DeepCopy` lose an unexported field? Write a test that proves it, then
describe a copy approach without that limitation.
