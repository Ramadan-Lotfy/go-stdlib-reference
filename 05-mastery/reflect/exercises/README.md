# reflect exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/reflect/exercises/...
```

## Exercise 1: StructToMap

Write `StructToMap(v any) map[string]any` that converts a struct (or a pointer
to one) into a map.

Rules:

- The key is the name from the `json` tag (the part before any comma), or the
  field name when there is no tag
- Skip unexported fields and fields tagged `json:"-"`
- Anything that is not a struct (including `nil`) returns `nil`

Hints: `reflect.ValueOf`, `Kind() == reflect.Pointer` then `Elem()`,
`t.Field(i).Tag.Get("json")`, `strings.Cut`, `IsExported()`.

## Exercise 2: MissingRequired

Write `MissingRequired(v any) []string` that returns the names of exported
fields tagged `validate:"required"` whose value is the zero value, in field order.

Hints: `Value.IsZero()`.

## Bonus
Write `Validate(v any) error` that also supports `validate:"min=3"` on strings
and ints. Why is parsing struct tags at run time slower than compile-time checks?
