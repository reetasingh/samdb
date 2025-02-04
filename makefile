# Define variables
GO = go
PKG = ./...  # This will run tests in the current directory and subdirectories

# Default target: run all tests
.PHONY: test
test:
	$(GO) test -v $(PKG)

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	$(GO) test -coverprofile=coverage.out $(PKG)
	$(GO) tool cover -html=coverage.out -o coverage.html

# Run tests with race detector
.PHONY: test-race
test-race:
	$(GO) test -race $(PKG)

# Run tests with a specific pattern (e.g., run only tests that match a name pattern)
.PHONY: test-pattern
test-pattern:
	$(GO) test -run <pattern> $(PKG)

# Clean up the build artifacts (optional)
.PHONY: clean
clean:
	$(GO) clean

# Run linters (optional, depends on whether you have linters like `golangci-lint` installed)
.PHONY: lint
lint:
	golangci-lint run

# Build the application
.PHONY: build
build:
	$(GO) build -o myapp

# Run the application
.PHONY: run
run:
	$(GO) run main.go

# Example for generating mocks for all interfaces
.PHONY: mock
mock:
	mockery --all  # This generates mocks for all interfaces in the package

# Example for generating mocks for all interfaces
.PHONY: pre-commit
pre-commit:
	pre-commit run --all-files
