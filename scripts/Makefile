.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

APP = gin-amis-admin
SERVER_BIN = ./${APP}

all: start

start:
	./${SERVER_BIN} web -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml --page ./configs/page_manager.yaml
