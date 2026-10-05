# embed exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/embed/exercises/...
```

## Exercise: embedded templates

The folder `solution/templates/` contains two files, `greeting.tmpl`
(`Hello, {{.Name}}!`) and `farewell.tmpl` (`Goodbye, {{.Name}}!`).

Write:

```go
func Render(name, who string) (string, error)
func Names() []string
```

- `Render("greeting", "Ada")` returns `Hello, Ada!`; an unknown name returns an error
- `Names()` returns `["farewell", "greeting"]`, derived from the embedded files
  (not hard-coded), so adding a template file adds a name

Hints: `//go:embed templates/*.tmpl` above an `embed.FS` variable,
`template.ParseFS`, `ExecuteTemplate`, `FS.ReadDir`.

## Bonus
Add a third template and watch `Names()` pick it up without code changes. Then
try to embed a file that does not exist and read the build error.
