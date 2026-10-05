# strconv exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/strconv/exercises/...
```

## Exercise 1: ParseKV

Write `ParseKV(s string) (map[string]int, error)` that parses `"a=1,b=2"`.

Rules:

- Pairs are separated by `,`, key and value by `=`
- Spaces around keys and values are ignored
- Empty input gives an empty map and no error
- A pair without `=` returns an error such as `invalid pair "b"`
- A bad number returns an error that **wraps** the strconv error, so
  `errors.Is(err, strconv.ErrSyntax)` is true. Include the key in the message:
  `invalid value for "a": strconv.Atoi: parsing "x": invalid syntax`

| Input | Result |
|-------|--------|
| `"a=1,b=2, c = 3"` | `map[a:1 b:2 c:3]` |
| `""` | `map[]` |
| `"a=1,b"` | error |
| `"a=x"` | error wrapping `ErrSyntax` |
| `"a=99999999999999999999"` | error wrapping `ErrRange` |

Hints: `strings.Cut`, `strconv.Atoi`, `fmt.Errorf` with `%w`.

## Exercise 2: HumanBytes

Write `HumanBytes(n int64) string` using 1024-based units (`B`, `KB`, `MB`,
`GB`, `TB`).

Rules:

- Below 1024 print the number and ` B`
- Otherwise print one decimal, drop a trailing `.0`, then the unit
- A value that would print as `1024.0 KB` must be shown in the next unit

| Input | Output |
|-------|--------|
| `0` | `0 B` |
| `1023` | `1023 B` |
| `1536` | `1.5 KB` |
| `1048575` | `1 MB` |
| `1500000` | `1.4 MB` |

Hints: `strconv.FormatInt`, `strconv.FormatFloat(v, 'f', 1, 64)`,
`strings.TrimSuffix`.

## Bonus

Add negative-number support to `HumanBytes`, then write `ParseHumanBytes` that
reverses it (`"1.5 KB"` gives `1536`) using `strconv.ParseFloat`.
