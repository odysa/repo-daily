.PHONY: build clean tool lint help

all: build

install:
	go install -v golang.org/x/lint/golint@latest
	go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
	go install -v github.com/kisielk/errcheck@latest

build:
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o telemetry-backend .

run:
	@go run .

check:
	@go fmt ./...
	@echo "go fmt done"
	@go vet ./...
	@echo "go vet done"
	@gocritic check ./...
	@echo "gocritic done"
	@errcheck -asserts -blank  ./...
	@echo "errcheck done"
	@golint ./...
	@echo "golint done"

test:
	@go test ./...

lint:
	@golint ./...


clean:
	@rm telemetry-backend
	@go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make fmt: format code"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"