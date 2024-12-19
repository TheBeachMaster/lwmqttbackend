.PHONY: clean build run

APP_NAME = mqttbackend
BUILD_DIR = $(PWD)/bin
OS = $(shell go env GOOS)
ARCH = $(shell go env GOARCH)

build:
	CGO_ENABLED=1 GOOS=$(OS) GOARCH=$(ARCH) go build --race -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

clean:
	rm -rf $(BUILD_DIR)

run:
	$(BUILD_DIR)/$(APP_NAME)
