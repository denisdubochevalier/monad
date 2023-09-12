all: 
	go build ./...

test:
	go test ./... -race -coverprofile=c.out -covermode=atomic

cover: test
	go tool cover -html=c.out
	
install:
	go install .

lint:
	golangci-lint run ./...
