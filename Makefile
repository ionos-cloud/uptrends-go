# GO related variables.
GO ?= go
GO_RUN_TOOLS ?= $(GO) run -modfile ./tools/go.mod
GO_TEST = $(GO_RUN_TOOLS) gotest.tools/gotestsum --format pkgname

.PHONY: generate
generate: ## Run generate commands.
	java -jar ./swagger-codegen-cli.jar \
    generate -DpackageName=uptrends \
    -i "https://api.uptrends.com/v4/swagger/v1/swagger.json" \
		-l go

.PHONY: fmt
fmt: ## Run go fmt against code.
	$(GO) run mvdan.cc/gofumpt -w .

.PHONY: vet
vet: ## Run go vet against code.
	$(GO) vet ./...
