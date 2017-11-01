GOFILES := $(shell find . -name "*.go")

.PHONY: check
check: test check-license

.PHONY: test
test:
	go test ./...

.PHONY: check-license
check-license:
	@echo "Checking for license in source files..."
	@grep -L LICENSE ${GOFILES}
	@exit `grep -L LICENSE ${GOFILES} | wc -l`

.PHONY: deps
deps:
	godeps -u dependencies.tsv

.PHONY: create-deps
create-deps:
	godeps -t ./... > dependencies.tsv
