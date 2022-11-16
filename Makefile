generate:
#	@go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f https://api.uptrends.com/v4/swagger/v1/swagger.json --skip-validation
	java -jar ./swagger-codegen-cli.jar \
    generate -DpackageName=uptrends \
    -i "https://api.uptrends.com/v4/swagger/v1/swagger.json" \
		-l go

.PHONY: fmt
fmt: ## Run go fmt against code.
	go run mvdan.cc/gofumpt -w .

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...
