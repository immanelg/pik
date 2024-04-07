default: build

build:
	go build -o pik ./app

fmt:
	go fmt ./...

test:
	go test -v ./...
