.PHONY: deps clean format build

all: clean format build

format:
	go fmt ./...

deps:
	go get -u ./...

clean: 
	rm -rf ./das-keyboard
	
build:
	go build -o das-keyboard ./cmd
