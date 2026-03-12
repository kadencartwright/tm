BINARY_NAME := tm
BUILD_DIR := ./bin
INSTALL_PATH := $(GOPATH)/bin

ifeq ($(GOPATH),)
	INSTALL_PATH := $(HOME)/go/bin
endif

.PHONY: all build test install clean

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Built $(BUILD_DIR)/$(BINARY_NAME)"

test:
	@echo "Running tests..."
	go test -v ./...

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@mkdir -p $(INSTALL_PATH)
	cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_PATH)/
	@echo "Installed to $(INSTALL_PATH)/$(BINARY_NAME)"

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned"
