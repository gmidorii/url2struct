BIN=u2s
MAIN_DIR=cmd/u2s

.PHONY: build
build:
	go build -o ./$(BIN) ./$(MAIN_DIR)/
