#!/usr/bin/env gmake
#
# TruthValidator Agent Makefile
#
# Provides build, test and deployment targets for:
# - Smart contracts
# - Go agent application
# - Local development environment
#
# Version: 1.0.0
# Maintainer: TruthValidator Team

# Include common build settings
include common.mk

# Phony targets (not actual files)
.PHONY: all contracts TruthValidatorAgent fmt tidy clean test install-deps run-localnet run-localnet-debug

# Default target - builds everything
all: contracts TruthValidatorAgent
	@printf "$(CYAN)*** Everything built successfully!$(OFF)\n"

# Build smart contracts
contracts:
	@printf "$(CYAN)*** Building $(BLUE)contracts$(CYAN)...$(OFF)\n"
	@$(MAKE) -C contracts

# Build main Go agent binary
TruthValidatorAgent:
	@printf "$(CYAN)*** Building $(BLUE)$@$(CYAN)...$(OFF)\n"
	@go build -o TruthValidatorAgent

# Format all Go code
fmt:
	@printf "$(CYAN)*** Formatting Go code...$(OFF)\n"
	@go fmt ./...

# Tidy Go module dependencies
tidy:
	@printf "$(CYAN)*** Tidying Go modules...$(OFF)\n"
	@go mod tidy

# Clean build artifacts
clean:
	@printf "$(CYAN)*** Cleaning up...$(OFF)\n"
	@go clean
	@$(MAKE) -C contracts clean

# Run standard end-to-end tests
test:
	@printf "$(CYAN)*** Running end-to-end tests...$(OFF)\n"
	@./test-e2e.sh

# Run tests against Sepolia testnet
test-sepolia-testnet:
	@printf "$(CYAN)*** Running end-to-end tests...$(OFF)\n"
	@./test-e2e-sepolia-testnet.sh

# Run tests against local Ethereum network
test-eth-localnet:
	@printf "$(CYAN)*** Running end-to-end tests...$(OFF)\n"
	@./test-e2e-eth-local.sh

# Install development dependencies
install-deps:
ifeq ($(shell which abigen),)
	@printf "$(CYAN)*** Installing dependency: $(BLUE)abigen$(CYAN)...$(OFF)\n"
	@go install github.com/ethereum/go-ethereum/cmd/abigen@latest
else
	@printf "$(CYAN)*** Dependency $(BLUE)abigen$(CYAN) is already installed.$(OFF)\n"
endif
ifeq ($(shell which solc),)
	@printf "$(CYAN)*** Installing dependency: $(BLUE)solc$(CYAN)...$(OFF)\n"
	@sudo snap install solc --edge
else
	@printf "$(CYAN)*** Dependency $(BLUE)solc$(CYAN) is already installed.$(OFF)\n"
endif

# Start local development network
run-localnet:
	@printf "$(CYAN)*** Starting $(BLUE)sapphire-localnet$(CYAN)...$(OFF)\n"
	@-docker run -it -p8545:8545 -p8546:8546 $(DOCKER_PLATFORM) ghcr.io/oasisprotocol/sapphire-localnet -test-mnemonic

# Start local development network in debug mode
run-localnet-debug:
	@printf "$(CYAN)*** Starting $(BLUE)sapphire-localnet$(CYAN) in $(MAGENTA)DEBUG$(CYAN) mode...$(OFF)\n"
	@-docker run -it -p8545:8545 -p8546:8546 $(DOCKER_PLATFORM) -e OASIS_NODE_LOG_LEVEL=debug -e LOG__LEVEL=debug ghcr.io/oasisprotocol/sapphire-localnet -test-mnemonic

# Help: Display available targets
help:
	@echo "Available targets:"
	@echo "  all              - Build everything (default)"
	@echo "  contracts        - Build smart contracts"
	@echo "  TruthValidatorAgent - Build main Go agent"
	@echo "  fmt              - Format Go code"
	@echo "  tidy             - Tidy Go modules"
	@echo "  clean            - Clean build artifacts"
	@echo "  test             - Run standard end-to-end tests"
	@echo "  test-sepolia-testnet - Run tests against Sepolia testnet"
	@echo "  test-eth-localnet - Run tests against local Ethereum network"
	@echo "  install-deps     - Install development dependencies"
	@echo "  run-localnet     - Start local development network"
	@echo "  run-localnet-debug - Start local network in debug mode"
