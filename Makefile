.PHONY: default build run fmt test clean

BUILDFLAGS := -trimpath -ldflags="-s -w"
NAME := "pik"

all: build

build:
	CGO_ENABLED=0 go build $(BUILDFLAGS) -o $(NAME)

run: build
	./$(NAME)

install:
	go install

fmt:
	go fmt ./...

test:
	go test -v ./...

clean:
	rm -f ./$(NAME)
	go clean
