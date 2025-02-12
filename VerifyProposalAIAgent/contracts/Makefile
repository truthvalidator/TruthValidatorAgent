#!/usr/bin/env gmake

include ../common.mk

.PHONY: all TruthValidatorSentientNet clean

all: TruthValidatorSentientNet
	@printf "$(CYAN)*** Contracts built successfully!$(OFF)\n"

TruthValidatorSentientNet:
	@printf "$(CYAN)*** Building contract: $(BLUE)TruthValidatorSentientNet$(CYAN)...$(OFF)\n"
	@$(MAKE) -C TruthValidatorSentientNet

clean:
	@printf "$(CYAN)*** Cleaning up...$(OFF)\n"
	@$(MAKE) -C TruthValidatorSentientNet clean

