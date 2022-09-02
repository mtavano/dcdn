#
# INTERNAL VARIABLES
#
BIN=$(PWD)/bin/
#
# TARGETS
#

run:
	echo "[run] Running node..."
	@export $$(cat .env) && go run cmd/node/main.go

dev:
	@echo "[dev] Running node in debug-hot-reload mode..."
	@export $$(cat .env) && nodemon --exec go run cmd/node/main.go --signal SIGTERM

build-arm:
	@echo "[build] Building arm node..."
	@GOOS=linux GOARCH=arm64 go build -o bin/node-arm cmd/node/main.go

build-linux:
	@echo "[build] Building linux node..."
	@GOOS=linux go build -o bin/node-linux cmd/node/main.go

build-macos:
	@echo "[build] Building macos node..."
	@GOOS=darwin go build -o bin/node-macos cmd/node/main.go

.PHONY: build-arm
