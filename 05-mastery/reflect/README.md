# reflect

## Overview
`reflect` lets a program inspect and modify values whose types are not known
at compile time: list struct fields and tags, call methods by name, build slices
and maps dynamically. It powers `encoding/json`, `fmt`, `database/sql` scanning
and validation libraries. It is slow, bypasses compile-time type checking and
panics easily, so use it in library-style code, never as the first tool.

## Core API

| Function / Type | Purpose |
|---|---|
| `reflect.TypeOf(x)`, `reflect.TypeFor[T]()` | The type: `Name()`, `Kind()`, `Elem()`, `NumField()`, `Field(i)` |
| `reflect.ValueOf(x)` | The value: `Kind()`, `Int()`, `String()`, `Field(i)`, `Len()`, `Index(i)`, `Interface()` |
| `Kind` (`Struct`, `Slice`, `Map`, `Pointer`, `Int`...) | The category of a type, used in `switch` |
| `StructField.Tag.Get("json")` | Read struct tags |
| `v.Elem()`, `v.CanSet()`, `v.SetInt`, `v.SetString` | Modify values (needs a pointer) |
| `v.MethodByName("X").Call(args)` | Call methods dynamically |
| `reflect.MakeSlice`, `MakeMap`, `New`, `Append` | Build values dynamically |
| `reflect.DeepEqual(a, b)` | Deep comparison (mostly for tests) |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/reflect/
```

## Gotchas
- **To modify a value you must pass a pointer** and call `.Elem()`: `reflect.ValueOf(&x).Elem()`. A value from `ValueOf(x)` is not addressable, so `Set*` panics.
- **Unexported fields cannot be read through `Interface()` or set.** Check `CanSet()` / `CanInterface()` / `IsExported()` first.
- **Type vs Kind:** a `type UserID int` has Kind `Int` but Name `UserID`. Switch on `Kind`, compare types with `==`.
- **Reflection panics instead of returning errors** (wrong kind, nil, unaddressable). Validate with `Kind()` before calling kind-specific methods.
- **`DeepEqual` treats a nil slice and an empty slice as different,** and does not ignore unexported fields. Prefer `slices.Equal` and `maps.Equal`, or `cmp.Diff` from go-cmp in tests.
- **It is slow** (often 10-100x a direct call) and allocates. Cache the results of type inspection in hot paths.
- **Generics are usually the better answer** for "work with many types" when you only need operations the types share.
- **`Pointer` vs `Ptr`:** `reflect.Ptr` is the old name of `reflect.Pointer`. Use `Pointer` in new code.

## Java comparison

| Java | Go |
|---|---|
| `Class<?>`, `getClass()` | `reflect.TypeOf(x)` |
| `getDeclaredFields()` | `t.Field(i)` for `i < t.NumField()` |
| annotations (`@NotNull`) | struct tags (`` `validate:"required"` ``) |
| `Method.invoke(obj, args)` | `v.MethodByName("X").Call(args)` |
| `Field.set(obj, v)` + `setAccessible` | `v.Elem().Field(i).Set(...)` (exported fields only) |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/reflect
- Source: https://cs.opensource.google/go/go/+/master:src/reflect/
