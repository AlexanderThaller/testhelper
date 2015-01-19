all:
	make format
	make test

format:
	find . -name "*.go" -not -path './Godeps/*' -type f -exec goimports -w=true {} \;

test:
	go test ./...
