# html/template

## Overview
`html/template` has the same API as `text/template` but understands HTML,
CSS, JavaScript and URL contexts and escapes values automatically for the place
they appear. This prevents most cross-site scripting (XSS) bugs without any
manual escaping.

## Core API

| Function / Type | Purpose |
|---|---|
| `template.New`, `Parse`, `ParseFiles`, `ParseFS`, `Must` | Same as `text/template` |
| `t.Execute`, `ExecuteTemplate` | Render with contextual escaping |
| `template.HTML`, `HTMLAttr`, `JS`, `URL` | Mark content as already safe (dangerous) |
| `template.HTMLEscapeString`, `JSEscapeString` | Manual escaping helpers |
| `{{define}}`, `{{template}}`, `{{block}}` | Layouts and partials |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/htmltemplate/
```

## Gotchas
- **Never convert user input to `template.HTML`, `template.JS` or `template.URL`.** Those types bypass escaping and reintroduce XSS.
- **Escaping depends on context:** the same value is escaped differently in text, an attribute, a URL or a script.
- **Do not build HTML with string concatenation and print it as a plain string.** It will be escaped again; pass data, not markup.
- **You cannot add or redefine templates after the first `Execute`.**
- **Parse templates once at startup**, not per request.
- **Errors can appear at execution time**, such as ambiguous contexts in script blocks. Test every template.
- **Use `ExecuteTemplate` with a named layout** to compose pages from partials.
- **HTML comments are stripped** from the template text.

## Java comparison

| Java | Go |
|---|---|
| Thymeleaf `th:text` (escaped) | `{{.}}` (escaped by default) |
| Thymeleaf `th:utext` (unescaped) | `template.HTML(...)` |
| JSP `<c:out>` | automatic contextual escaping |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/html/template
- Source: https://cs.opensource.google/go/go/+/master:src/html/template/
