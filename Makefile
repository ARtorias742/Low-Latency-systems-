# Makefile for low-latency-chat project

BINARY_SERVER = low-latency-chat-server
BINARY_CLIENT = low-latency-chat-client
CMD_SERVER = cmd/server/main.go
CMD_CLIENT = cmd/client/main.go

.PHONY: all
all: build

.PHONY: build
build: build-server build-client

.PHONY: build-server
build-server:
	@echo "Building server..."
	go build -o $(BINARY_SERVER) $(CMD_SERVER)

.PHONY: build-client
build-client:
	@echo "Building client..."
	go build -o $(BINARY_CLIENT) $(CMD_CLIENT)

.PHONY: run-server
run-server:
	@echo "Starting server on :8080..."
	go run $(CMD_SERVER)

.PHONY: run-client
run-client:
	@if [ -z "$(USER)" ]; then \
		echo "Error: USER variable is required. Usage: make run-client USER=<username>"; \
		exit 1; \
	else \
		echo "Starting client as $(USER)..."; \
		go run $(CMD_CLIENT) $(USER); \
	fi

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_SERVER) $(BINARY_CLIENT)

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

.PHONY: vet
vet:
	@echo "Vetting code..."
	go vet ./...

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build         - Build server and client"
	@echo "  make build-server  - Build server only"
	@echo "  make build-client  - Build client only"
	@echo "  make run-server    - Run the server"
	@echo "  make run-client USER=<username> - Run a client (e.g., make run-client USER=user1)"
	@echo "  make clean         - Remove built binaries"
	@echo "  make fmt           - Format the code"
	@echo "  make vet           - Vet the code for issues"
	@echo "  make help          - Show this help message"