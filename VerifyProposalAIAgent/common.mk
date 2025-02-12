#!/usr/bin/env gmake

# Check if we're running in an interactive terminal.
ISATTY := $(shell [ -t 0 ] && echo 1)

ifdef ISATTY
# Running in interactive terminal, OK to use colors!
GREEN = \e[32;1m
BLUE = \e[34;1m
MAGENTA = \e[35;1m
CYAN = \e[36;1m
OFF = \e[0m
else
# Don't use colors if not running interactively.
GREEN = ""
BLUE = ""
MAGENTA = ""
CYAN = ""
OFF = ""
endif

# Check if we're running on an x86_64.
MACHINE := $(shell uname -m)
ifneq ($(MACHINE),x86_64)
DOCKER_PLATFORM = --platform linux/x86_64
else
DOCKER_PLATFORM =
endif

