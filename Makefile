install:
	dep ensure

build:
	go build ./...

test:
    go test ./...