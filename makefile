.PHONY: build test run

APP_MAIN_FILE := ./cmd/main.go

build:
	@go build -o bin/design-patterns $(APP_MAIN_FILE)

test: 
	@go test -cover -v ./...

run: build  
	@./bin/design-patterns
