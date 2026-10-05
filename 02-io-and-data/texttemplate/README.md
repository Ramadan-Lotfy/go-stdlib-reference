# text/template

## Overview
`text/template` fills a text template with data. Actions in `{{ }}` print
values, loop with `range`, branch with `if`, call functions and include other
templates. It is used for config files, code generation, emails and CLI output.
For HTML, use `html/template` instead.

## Core API

| Function / Type | Purpose |
|---|---|
| `template.New(name).Parse(text)` | Create and parse a template; `template.Must` panics on error |
| `t.Execute(w, data)`, `t.ExecuteTemplate(w, name, data)` | Render |
| `{{.Field}}`, `{{.Key}}`, `{{$x := ...}}` | Fields, map keys, variables |
| `{{range}}...{{else}}...{{end}}`, `{{if}}`, `{{with}}` | Control flow |
| `{{define "x"}}`, `{{template "x" .}}`, `{{block}}` | Reusable and overridable parts |
| `{{- ... -}}` | Trim surrounding whitespace |
| `template.FuncMap` | Add your own functions |
| `eq`, `ne`, `lt`, `gt`, `and`, `or`, `not`, `len`, `index`, `printf` | Built-in functions |
| `Option("missingkey=error")` | Fail on missing map keys |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/texttemplate/
```

## Gotchas
- **Use `html/template` for anything that ends up in a web page.** `text/template` does no escaping.
- **The dot (`.`) changes inside `range` and `with`.** Use `$` for the root data.
- **Parse once, execute many times.** A parsed template is safe for concurrent `Execute` calls.
- **Whitespace around actions is kept.** Use `{{-` and `-}}` to trim it.
- **Missing map keys print `<no value>` by default.** Set `missingkey=error` to fail fast.
- **Functions must be registered before `Parse`**, with `Funcs`.
- **Comparison functions are strict about types:** `eq .Age 3` fails if `Age` is a `float64`. Convert in Go.
- **Methods and exported fields are available; unexported ones are not.**
- **Execution errors may happen after partial output has been written.** Render to a buffer first if that matters.

## Java comparison

| Java | Go |
|---|---|
| Thymeleaf / FreeMarker / Mustache | `text/template` (text) or `html/template` (HTML) |
| `${user.name}` | `{{.User.Name}}` |
| `th:each` | `{{range}}` |
| `th:if` | `{{if}}` |
| `th:fragment` + `th:replace` | `{{define}}` + `{{template}}` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/text/template
- Source: https://cs.opensource.google/go/go/+/master:src/text/template/
