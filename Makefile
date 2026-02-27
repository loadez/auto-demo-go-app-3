build:
	go build -o bin/app ./...

test:
	go test -race -v ./...

lint:
	golangci-lint run ./...

run: build
	./bin/app
