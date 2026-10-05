# strings exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/strings/exercises/...
```

## Exercise 1: Slugify

Write `Slugify(s string) string` that turns a title into a URL slug.

Rules:

- Keep letters and digits, lowercased
- Every run of any other characters becomes a single `-`
- No leading or trailing `-`

| Input | Output |
|-------|--------|
| `"Hello, World!"` | `"hello-world"` |
| `"  Go   1.22 Released  "` | `"go-1-22-released"` |
| `"Spring Boot & Go"` | `"spring-boot-go"` |
| `"---"` | `""` |

Hints: `strings.Map`, `strings.Fields`, `strings.Join`, `unicode.IsLetter`.

## Exercise 2: Word counter

Write `CountWords(text string) map[string]int`.

Rules:

- Words are case-insensitive (`The` and `the` are the same word)
- Ignore these characters around a word: `. , ; : ! ? " ' ( )`
- Tokens that are only punctuation are skipped

Example: `The cat and the hat. The "end"!` gives
`map[the:3 cat:1 and:1 hat:1 end:1]`.

Hints: `strings.ToLower`, `strings.Fields`, `strings.Trim`.

## Bonus

Make `Slugify` ASCII-only by transliterating nothing and dropping non-ASCII
letters. Which `strings` and `unicode` functions change?
