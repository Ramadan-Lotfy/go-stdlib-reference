# Contributing

Thanks for helping make this a great reference. Keep every package consistent.

## Adding a package

```bash
make new STAGE=01-foundations PKG=fmt
```

This creates `01-foundations/fmt/` from `templates/package/`. For nested stdlib packages,
drop the slash (for example `PKG=nethttp` for `net/http`).

Then fill in:

1. **README.md**: follow the section order in the template.
2. **example_test.go**: write `Example` functions with an `// Output:` comment so
   `go test` verifies them.
3. **exercises/**: at least one exercise with a solution.

## Package README sections (in order)

1. Overview: what it is for, in two sentences
2. Core API: the 5 to 10 functions that matter
3. Examples: point to `example_test.go`
4. Gotchas: common mistakes
5. Java comparison: the closest Java equivalent (optional)
6. Exercises
7. Links: pkg.go.dev and source

## Style rules

- Code must pass `gofmt` and `go vet`.
- Prefer small, focused examples over long programs.
- Every example needs an `// Output:` comment unless output is non-deterministic.
- Name examples `Example_topic` (for example `Example_numbers`). The reference packages
  define no functions of their own, so avoid `ExamplePrintf`-style names that refer
  to identifiers `go vet` cannot find.
- Keep prose short and concrete.

## Before opening a PR

```bash
make check
```
