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
	go build -o ./bin/client ./src/cmd
	go build -o ./bin/server ./src/server
	go build -o ./bin/editor ./src/editor
