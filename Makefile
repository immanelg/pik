BINARY := "pik"

default: build

build:
	go build -o $(BINARY) ./app

run: build
	./$(BINARY)

fmt:
	go fmt ./...

test:
	go test -v ./...
