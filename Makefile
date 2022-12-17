compile:
	go build -o build/ cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./...

clean:
	rm -r build
