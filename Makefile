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

bench:
	go test --cover -bench . -benchmem ./...

setup-ci:
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls

test-ci: vet setup-ci
	go test -v -covermode count -coverprofile coverage.out
	go tool cover -html cover.out -o cover.html
	$HOME/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN
