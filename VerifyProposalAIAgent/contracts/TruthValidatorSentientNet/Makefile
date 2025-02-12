#!/usr/bin/env gmake

include ../../common.mk

.PHONY: TruthValidatorSentientNet clean

TruthValidatorSentientNet:
	@printf "$(CYAN)*** Compiling $(BLUE)$@$(CYAN)...$(OFF)\n"
	@solc --bin --abi --optimize --overwrite -o . $@.sol
	@mv $@.bin $@.hex
	@printf "$(CYAN)*** Generating Go bindings for $(BLUE)$@$(CYAN)...$(OFF)\n"
	@abigen --abi=$@.abi --bin=$@.hex --pkg TruthValidatorSentientNet --out=$@.go

clean:
	@printf "$(CYAN)*** Cleaning up...$(OFF)\n"
	@-rm -f TruthValidatorSentientNet.abi TruthValidatorSentientNet.hex TruthValidatorSentientNet.go

