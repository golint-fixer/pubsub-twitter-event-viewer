
# Go parameters
BINARY_NAME=tview

all: test

deps:
	go get -u github.com/tools/godep
	godep restore

build:
	go build -v -o ./bin/$(BINARY_NAME)

test:
	go test -v ./...

clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)

run: build
	open http://127.0.0.1:8080
	bin/$(BINARY_NAME)
