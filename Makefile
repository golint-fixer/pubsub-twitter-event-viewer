
# Go parameters
BINARY_NAME=tview
GCP_PROJECT_NAME=knative-samples

all: test

deps:
	go get github.com/golang/dep/cmd/dep
	dep ensure

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

gcr:
	gcloud container builds submit --project=$(GCP_PROJECT_NAME) --tag gcr.io/$(GCP_PROJECT_NAME)/$(BINARY_NAME):latest .


