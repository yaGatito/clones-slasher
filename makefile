.PHONY: build-win build-lin clean

BINARY_NAME := cloneseeker
OUTPUT_DIR := out

tools:
	irm get.scoop.sh | iex
	scoop install pwsh

build-win-amd64:
	mkdir .\$(OUTPUT_DIR)\win-amd64
	pwsh -Command '$$env:GOOS="windows"; $$env:GOARCH="amd64"; go build -ldflags="-s -w" -o ./out/win-amd64/cloneseeker.exe ./cmd/main.go'

### NOT TESTED
build-lin-amd64:
	mkdir -p ./$(OUTPUT_DIR)/lin-amd64
	pwsh -Command '$$env:CGO_ENABLED="0"; $$env:GOOS="linux"; $$env:GOARCH="amd64"; go build -ldflags="-s -w" -o ./out/lin-amd64/cloneseeker.exe ./cmd/app/main.go'

### TODO
clean:
	rm $(OUTPUT_DIR)

all: build-win build-lin
