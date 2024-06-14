
default: build

build:
	go build

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
