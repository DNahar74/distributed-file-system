.PHONY: build run test

build:
	@echo "🔨 Building..."
	go build -o bin/fileStore.exe

run: build
	@echo "🚀 Running..."
	@./bin/fileStore.exe

test:
	@echo "🧪 Running tests..."
	@go test ./... -v