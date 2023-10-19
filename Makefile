all: test

test:
	go test -v ./...

build:
	go build

install:
	go install

clean:
	rm -f mercury

sender:
	cd examples/sender && go run .

receiver:
	cd examples/receiver && go run .

.PHONY: all test build install clean sender receiver
