# Makefile for low-latency-chat project

# Variables
BINARY_NAME = low-latency-chat
CMD_DIR = cmd
MAIN_FILE = $(CMD_DIR)/main.go

# Default target
.PHONY: all
all: build

# Initialize Go module (run once)
.PHONY: init
init:
	go mod init low-latency-chat

# Tidy dependencies
.PHONY: tidy
tidy:
	go mod tidy

# Build the project
.PHONY: build
build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# Run the project
.PHONY: run
run:
	go run $(MAIN_FILE)

# Build and run
.PHONY: build-run
build-run: build
	./$(BINARY_NAME)

# Clean up generated files
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Vet the code for potential issues
.PHONY: vet
vet:
	go vet ./...

# Test (if tests are added later)
.PHONY: test
test:
	go test ./...

# Help command to list available targets
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make init      - Initialize Go module (run once)"
	@echo "  make tidy      - Tidy Go module dependencies"
	@echo "  make build     - Build the project"
	@echo "  make run       - Run the project without building"
	@echo "  make build-run - Build and run the project"
	@echo "  make clean     - Remove built binary"
	@echo "  make fmt       - Format the code"
	@echo "  make vet       - Vet the code for issues"
	@echo "  make test      - Run tests (if any)"
	@echo "  make help      - Show this help message"