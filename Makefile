# Makefile for Go project

# Go commands
GO      = go
GOFMT   = gofmt
GOTEST  = go test
GOBUILD = go build
GORUN   = go run

# Directories
SRC_DIR = .
BIN_DIR = bin

# Executable name
EXE = $(BIN_DIR)/main

# Source files
SRC = $(wildcard $(SRC_DIR)/*.go)

# Default target
.DEFAULT_GOAL := all

# Build the binary
.PHONY: all
all: $(EXE)

$(EXE): $(SRC)
	@echo "Building binary..."
	@mkdir -p $(BIN_DIR)
	$(GOBUILD) -o $(EXE) $(SRC)

# Format the source code
.PHONY: fmt
fmt:
	@echo "Formatting source code..."
	$(GOFMT) -w $(SRC_DIR)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) ./...

# Clean build files
.PHONY: clean
clean:
	@echo "Cleaning build files..."
	@rm -rf $(BIN_DIR)

# Run the built executable
.PHONY: run
run: $(EXE)
	@echo "Running binary..."
	./$(EXE)

# Run the Go source files directly
.PHONY: go_run
go_run:
	@echo "Running source files..."
	$(GORUN) $(SRC)
