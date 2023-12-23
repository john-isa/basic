init: distclean tidy

build:
	go build -o build/basic cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./...

lint:
	golangci-lint run ./...

clean:
	if test -d ./build; then rm -r build; fi

distclean: clean
	test -d "./vendor" && rm -r vendor
	rm go.sum

update: update_libs tidy

update_libs:
	go get -t -u ./...

tidy:
	go mod tidy
	go mod vendor
