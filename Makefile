.PHONY: all
coverage:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"
