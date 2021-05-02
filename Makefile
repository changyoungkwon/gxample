
.PHONY: download-deps
download-deps: ## download dependencies
	@echo Download go.sum dependencies
	go mod download

.PHONY: install-tools
install-tools: download-deps
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: build
build: | fmt lint ## fmt, lint, test and build
	go build -o ./out/gxample ./cmd/gxample

.PHONY: test
test: ## run test
	go test -race -cover ./...

.PHONY: fmt
fmt: ## run format
	test -z $(shell go fmt ./...)

.PHONY: clean
clean: ## clean build artifacts
	@rm -rf ./out || true

.PHONY: lint
lint: ## run golint
	golint -set_exit_status ./...
