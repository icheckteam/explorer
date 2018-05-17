PACKAGES=$(shell go list ./... | grep -v '/vendor/')
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_FLAGS = -ldflags "-X github.com/icheckteam/explorer/version.GitCommit=${COMMIT_HASH}"

get_vendor_deps:
	@rm -rf vendor/
	@echo "--> Running dep ensure"
	@dep ensure -v



########################################
### Build
# This can be unified later, here for easy demos
build:
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/explorercli.exe ./cmd/main.go
else
	go build $(BUILD_FLAGS) -o build/explorercli ./cmd/main.go
endif
