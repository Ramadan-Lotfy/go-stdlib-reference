.PHONY: test check fmt vet new

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

check:
	@test -z "$$(gofmt -l .)" || (echo "gofmt needed: run 'make fmt'" && exit 1)
	go vet ./...
	go test -race ./...

# Usage: make new STAGE=01-foundations PKG=fmt
new:
	@test -n "$(STAGE)" -a -n "$(PKG)" || (echo "usage: make new STAGE=01-foundations PKG=fmt" && exit 1)
	@test ! -d "$(STAGE)/$(PKG)" || (echo "$(STAGE)/$(PKG) already exists" && exit 1)
	@mkdir -p $(STAGE)/$(PKG)/exercises
	@touch $(STAGE)/$(PKG)/exercises/.gitkeep
	@sed 's/__PKG__/$(PKG)/g; s/__NAME__/$(PKG)ref/g' templates/package/README.md.tmpl > $(STAGE)/$(PKG)/README.md
	@sed 's/__PKG__/$(PKG)/g; s/__NAME__/$(PKG)ref/g' templates/package/doc.go.tmpl > $(STAGE)/$(PKG)/doc.go
	@sed 's/__PKG__/$(PKG)/g; s/__NAME__/$(PKG)ref/g' templates/package/example_test.go.tmpl > $(STAGE)/$(PKG)/example_test.go
	@echo "created $(STAGE)/$(PKG)"
