.PHONY: download-deps
download-deps: ## download dependencies
	@echo Download go.sum dependencies
	go mod download

.PHONY: install-tools
install-tools: download-deps
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: gendoc
gendoc: | install-tools
	swag init -g ./internal/routes/routes.go --parseInternal

.PHONY: build
build: | fmt lint ## fmt, lint, test and build
	go build -o ./out/gxample ./cmd/gxample

.PHONY: run
run: | build
	./out/gxample serve

.PHONY: test
test: ## run test
	go test -race -cover ./...

.PHONY: fmt
fmt: ## run format
	test -z $(shell go fmt ./...)

.PHONY: clean
clean: ## clean build artifacts
	@rm -rf ./out || true
	go clean -modcache
	go clean -cache

.PHONY: reset-db
reset-db: ## reset database including migration
	@PGPASSWORD=${POSTGRES_PASSWORD} psql -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_DB} -a -f scripts/reset_db.sql 
.PHONY: dump-db
dump-db: ## dump database
	@PGPASSWORD=${POSTGRES_PASSWORD} pg_dump -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} ${POSTGRES_DB} > dump.sql 
	

.PHONY: lint
lint: ## run golint
	golint -set_exit_status ./...
