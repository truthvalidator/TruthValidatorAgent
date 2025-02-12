#!/usr/bin/env gmake

include common.mk

.PHONY: all contracts TruthValidatorAgent fmt tidy clean test install-deps run-localnet run-localnet-debug

all: contracts TruthValidatorAgent
	@printf "$(CYAN)*** Everything built successfully!$(OFF)\n"

contracts:
	@printf "$(CYAN)*** Building $(BLUE)contracts$(CYAN)...$(OFF)\n"
	@$(MAKE) -C contracts

TruthValidatorAgent:
	@printf "$(CYAN)*** Building $(BLUE)$@$(CYAN)...$(OFF)\n"
	@go build -o TruthValidatorAgent

fmt:
	@printf "$(CYAN)*** Formatting Go code...$(OFF)\n"
	@go fmt ./...

tidy:
	@printf "$(CYAN)*** Tidying Go modules...$(OFF)\n"
	@go mod tidy

clean:
	@printf "$(CYAN)*** Cleaning up...$(OFF)\n"
	@go clean
	@$(MAKE) -C contracts clean

test:
	@printf "$(CYAN)*** Running end-to-end tests...$(OFF)\n"
	@./test-e2e.sh

test-sepolia-testnet:
	@printf "$(CYAN)*** Running end-to-end tests...$(OFF)\n"
	@./test-e2e-sepolia-testnet.sh

test-eth-localnet:
	@printf "$(CYAN)*** Running end-to-end tests...$(OFF)\n"
	@./test-e2e-eth-local.sh

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

run-localnet:
	@printf "$(CYAN)*** Starting $(BLUE)sapphire-localnet$(CYAN)...$(OFF)\n"
	@-docker run -it -p8545:8545 -p8546:8546 $(DOCKER_PLATFORM) ghcr.io/oasisprotocol/sapphire-localnet -test-mnemonic

run-localnet-debug:
	@printf "$(CYAN)*** Starting $(BLUE)sapphire-localnet$(CYAN) in $(MAGENTA)DEBUG$(CYAN) mode...$(OFF)\n"
	@-docker run -it -p8545:8545 -p8546:8546 $(DOCKER_PLATFORM) -e OASIS_NODE_LOG_LEVEL=debug -e LOG__LEVEL=debug ghcr.io/oasisprotocol/sapphire-localnet -test-mnemonic
