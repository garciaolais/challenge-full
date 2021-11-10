BIN = 'server'
VERSION	?= $(shell go run cmd/main.go -version)

.PHONY: test fmt build test version

clean:	
	go clean -i ./...
	rm -f $(BIN)
fmt:
	go fmt ./...
build:
	go build -o $(BIN) cmd/main.go 
test:
	go test -v test/*.go -short
docker-build:
	docker build -t $(BIN):$(VERSION) .
