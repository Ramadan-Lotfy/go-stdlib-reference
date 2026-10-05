# path/filepath exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/filepath/exercises/...
```

## Exercise 1: SafeJoin

Write `SafeJoin(base, userPath string) (string, error)` that joins an
untrusted relative path onto `base` but returns an error if the path could
escape `base`.

| Input path | Result |
|------------|--------|
| `a/b.txt` | `base/a/b.txt` |
| `../etc/passwd` | error |
| `/etc/passwd` | error |
| `a/../../b` | error |
| `""` | error |

Hints: `filepath.IsLocal` checks lexically; call it **before** `filepath.Join`.

## Exercise 2: ChangeExt

Write `ChangeExt(name, newExt string) string` that swaps the last extension
(`report.txt` and `.md` gives `report.md`). A name with no extension just gets
`newExt` added.

Hints: `filepath.Ext`, `strings.TrimSuffix`.

## Bonus
`IsLocal` is lexical only. What attack does it not prevent? (Hint: symlinks. Read about `os.Root` in newer Go versions.)
