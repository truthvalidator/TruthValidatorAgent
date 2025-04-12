#!/usr/bin/env gmake
#
# Common Makefile Definitions
#
# Provides shared variables and functions used across all Makefiles
# - Terminal color definitions
# - Platform detection
# - Common build settings
#
# Version: 1.0.0
# Maintainer: TruthValidator Team

# ================= Terminal Colors =================
# Detect if running in interactive terminal (tty)
ISATTY := $(shell [ -t 0 ] && echo 1)

ifdef ISATTY
# Color definitions for interactive terminals
GREEN = \e[32;1m   # Success messages
BLUE = \e[34;1m    # Informational messages
MAGENTA = \e[35;1m # Warning messages
CYAN = \e[36;1m    # Highlighted text
OFF = \e[0m        # Reset color
else
# Empty color definitions for non-interactive use
GREEN = ""
BLUE = ""
MAGENTA = ""
CYAN = ""
OFF = ""
endif

# ================= Platform Detection =================
# Detect CPU architecture for Docker platform compatibility
MACHINE := $(shell uname -m)
ifneq ($(MACHINE),x86_64)
# Force x86_64 platform for non-x86_64 hosts
DOCKER_PLATFORM = --platform linux/x86_64
else
# Native platform for x86_64 hosts
DOCKER_PLATFORM =
endif

# ================= Usage Example =================
# To use these variables in other Makefiles:
# @printf "$(GREEN)Success!$(OFF)\n"
# @printf "$(BLUE)Processing...$(OFF)\n"
