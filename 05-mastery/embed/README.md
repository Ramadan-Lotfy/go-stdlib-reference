# embed

## Overview
`embed` compiles files into the program binary, so a single executable can
carry its own templates, SQL migrations, static web assets and default config.
You add a `//go:embed` comment directly above a variable of type `string`,
`[]byte` or `embed.FS`. `embed.FS` implements `fs.FS`, so it works with
`html/template`, `http.FileServer`, `fs.WalkDir` and everything else that takes a
file system.

## Core API

| Function / Type | Purpose |
|---|---|
| `//go:embed path` | Directive above a package-level variable (`string`, `[]byte` or `embed.FS`) |
| `//go:embed dir other/*.css` | Several patterns; directories are embedded recursively |
| `embed.FS` | Read-only file system: `Open`, `ReadFile`, `ReadDir` |
| `fs.Sub(fsys, "dir")` | Re-root an embedded tree (for example to serve `static/` at `/`) |
| `template.ParseFS(fsys, "*.tmpl")` | Load templates from embedded files |
| `http.FileServer(http.FS(fsys))` | Serve embedded assets |
| `all:` prefix | Include files starting with `.` or `_` |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/embed/
```

## Gotchas
- **The directive must be directly above a package-level `var`** and the file must `import "embed"` (or `_ "embed"` for `string` and `[]byte`).
- **Paths are relative to the Go source file's directory** and cannot contain `..`, cannot start with `/`, and cannot reach into another module.
- **Files whose names start with `.` or `_` are skipped** when embedding a directory. Use `//go:embed all:dir` to include them.
- **Embedded data is read-only and baked in at build time.** Changing the file requires a rebuild. It also grows the binary, so do not embed gigabytes.
- **Names inside an `embed.FS` keep the directory prefix.** `//go:embed static/app.css` is read as `static/app.css`; use `fs.Sub` to drop the prefix.
- **A missing file is a build error,** not a run-time one. That is a feature: you cannot ship a binary without its templates.
- **`string` and `[]byte` variables hold the whole file** and the `[]byte` is a copy you may modify; the `string` is immutable.
- **During development you often want live reload.** Use `os.DirFS` behind a flag in dev builds and `embed.FS` in release builds; both are an `fs.FS`.

## Java comparison

| Java | Go |
|---|---|
| `classpath:` resources / `getResourceAsStream` | `//go:embed` + `embed.FS` |
| Spring Boot fat JAR with static assets | single binary with embedded assets |
| Flyway `db/migration` scripts on the classpath | embedded `*.sql` files |
| Thymeleaf `templates/` directory | embedded `templates/*.tmpl` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/embed
- Source: https://cs.opensource.google/go/go/+/master:src/embed/
