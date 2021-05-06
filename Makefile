.DEFAULT_GOAL := buildq

fmt:
	go fmt ./...

lint: fmt
	golint ./...

vet: fmt
	go vet ./...

run:
	go run .

test:
	go test --cover

test-ci: vet
	go test --cover

bench:
	go test --cover -bench . -benchmem ./...
