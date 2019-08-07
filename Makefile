# Makefile for Go POSIX

PROGRAM := posix
DEPS := $(shell find . -name '*.go')
V ?= 1

ifeq ($(V),0)
Q := @
else
Q :=
endif

all: symlinks
.PHONY: symlinks
symlinks: $(PROGRAM)
	$(Q)find ./cmds -mindepth 1 -maxdepth 1 -type d -exec basename {} \; |\
		xargs -n1 ln -snf $^

$(PROGRAM): $(DEPS)
	@echo "Building $@"
	$(Q)go build .

.PHONY: clean
clean:
	@echo "Deleting symlinks"
	$(Q)find . -maxdepth 1 -type l -delete
	@echo "Deleting $(PROGRAM)"
	$(Q)rm -f $(PROGRAM)

.PHONY: format
format:
	@echo "Running go fmt"
	$(Q)go fmt .
	$(Q)go fmt ./cli
	$(Q)go fmt ./cmds/*
