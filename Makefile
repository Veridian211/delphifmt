build:
	go build -o ./bin/delphifmt ./delphifmt/main.go

run:
	go run ./dev/main.go

run-%:
	go run ./dev/main.go $*

test:
	go test ./...
