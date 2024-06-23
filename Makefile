.PHONY: default build run fmt test clean

default: build

build:
	go build -o pik

run: build
	./pik

install:
	go install

fmt:
	go fmt ./...

test:
	go test -v ./...

clean:
	rm -f ./pik
	go clean
