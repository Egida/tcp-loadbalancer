SERVER_DIR = server
BIN_DIR = bin
CMD_DIR = cmd
SHELL := bash
CURENT_DIR := $(shell pwd)
SHELL_VERSION := $(shell bash --version | head -1)
ifeq ($(origin $(SHELL_VERSION)), undefined)
	SHELL := sh
endif

server: $@
	@echo "build server"
	go build  -o ${BIN_DIR}/$@ ${CURENT_DIR}/${CMD_DIR}/$@/
	chmod +xwr ${BIN_DIR}/$@

run-server: server
	${BIN_DIR}/server

## create serever and run
run: server
	${BIN_DIR}/server
.PHONY: server
