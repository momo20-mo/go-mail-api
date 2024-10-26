.PHONY: all build test lint security doc clean help

# Define variables
GO := go
GOLINT := golint
GOLANGCI_LINT := golangci-lint
GODOC := godoc

# Check for necessary tools
ifeq (, $(shell which $(GO)))
	$(error "Go not found. Please install Go")
endif

# Default target
all: build test lint

# Build the Go application
build:
	@echo "Building Go application..."
	@$(GO) build -o dist/app cmd/main.go

# Run tests
test:
	@echo "Running tests..."
	@$(GO) test ./... -v

# Run static analysis with golint and golangci-lint
lint:
	@if command -v $(GOLANGCI_LINT) >/dev/null 2>&1; then \
		echo "Running golangci-lint..."; \
		$(GOLANGCI_LINT) run; \
	elif command -v $(GOLINT) >/dev/null 2>&1; then \
		echo "Running golint..."; \
		$(GOLINT) ./...; \
	else \
		echo "Neither golangci-lint nor golint found. Please install one to run static code analysis."; \
		exit 1; \
	fi

# Run security analysis (if gosec is available)
security:
	@if command -v gosec >/dev/null 2>&1; then \
		echo "Running gosec for security analysis..."; \
		gosec ./...; \
	else \
		echo "gosec not found. Please install gosec to run security analysis."; \
		exit 1; \
	fi

# Generate Go documentation (requires godoc)
doc:
	@if command -v $(GODOC) >/dev/null 2>&1; then \
		echo "Starting documentation server on http://localhost:6060..."; \
		$(GODOC) -http=:6060; \
	else \
		echo "godoc not found. Please install godoc to generate documentation."; \
		exit 1; \
	fi

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf bin
	@echo "Cleaned up build artifacts."

# Help
help:
	@echo "Available targets:"
	@echo "  build    - Build the Go application"
	@echo "  test     - Run tests for the Go application"
	@echo "  lint     - Run linting and static analysis (requires golint or golangci-lint)"
	@echo "  security - Run security analysis (requires gosec)"
	@echo "  doc      - Start documentation server (requires godoc)"
	@echo "  clean    - Remove build artifacts"
	@echo ""
	@echo "Usage: make [target]"
