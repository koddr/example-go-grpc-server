.PHONY: build

build:
	@go build -o ./build/grpc_server ./cmd/server/...

.PHONY: run

run: build
	@./build/grpc_server