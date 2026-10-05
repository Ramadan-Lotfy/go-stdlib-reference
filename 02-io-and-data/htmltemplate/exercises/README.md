# html/template exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/htmltemplate/exercises/...
```

## Exercise: RenderComments

Write `RenderComments(comments []string) (string, error)` that renders a list
of user comments as `<ul><li>...</li></ul>` using `html/template`.

```
["nice", "<script>x</script>"]
-> <ul><li>nice</li><li>&lt;script&gt;x&lt;/script&gt;</li></ul>
```

Rules:

- Markup inside a comment must come out escaped
- No comments gives `<ul></ul>`
- Parse the template once, at package level

Hints: `template.Must(template.New(...).Parse(...))`, `range`.

## Bonus
Switch the import to `text/template` and run the test. What breaks, and why
does that matter for a comments section?
