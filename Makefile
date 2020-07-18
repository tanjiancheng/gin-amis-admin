.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

APP = gin-amis-admin
SERVER_BIN = ./cmd/${APP}/${APP}
RELEASE_ROOT = release
RELEASE_SERVER = release/${APP}

all: start

build:
	@go build -ldflags "-w -s" -o $(SERVER_BIN) ./cmd/${APP}

build-darwin:
	xgo -go go-1.14.x -targets=darwin/amd64 -pkg=cmd/gin-amis-admin/main.go -dest=cmd/${APP} -out=gin-amis-admin .

start:
	go run cmd/${APP}/main.go web -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml --page ./configs/page_manager.yaml

swagger:
	swag init --generalInfo ./internal/app/swagger.go --output ./internal/app/swagger

wire:
	wire gen ./internal/app/injector

test:
	@go test -v $(shell go list ./...)

clean:
	rm -rf data release $(SERVER_BIN) ./internal/app/test/data ./cmd/${APP}/data

pack: build
	mkdir -p $(RELEASE_SERVER)
	rm -f $(APP)-linux-adm64.tar.gz
	cp -r $(SERVER_BIN) configs web $(RELEASE_SERVER)
	cp scripts/pack/* $(RELEASE_SERVER)
	cd $(RELEASE_ROOT) && tar -zcvf $(APP)-linux-adm64.tar.gz ${APP} && sudo rm -rf ${APP}
pack-darwin: build-darwin
	mkdir -p $(RELEASE_SERVER)
	rm -f $(APP)-darwin-amd64.tar.gz
	cp -r $(SERVER_BIN)-darwin-10.6-amd64 configs web $(RELEASE_SERVER)
	cp scripts/pack/* $(RELEASE_SERVER)
	mv $(RELEASE_SERVER)/$(APP)-darwin-10.6-amd64 $(RELEASE_SERVER)/$(APP)
	cd $(RELEASE_ROOT) && tar -zcvf $(APP)-darwin-amd64.tar.gz ${APP} && sudo rm -rf ${APP}
