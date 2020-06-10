.PHONY: deps clean format build

all: clean format build

format:
	go fmt ./...

deps:
	go get -u ./...

clean: 
	rm -rf ./bin
	
build:
	mkdir -p ./bin/
	go build -o ./bin/client ./cmd
	go build -o ./bin/server ./server
