$(shell mkdir -p artifacts)

build: clean
	goreleaser build --snapshot --clean

lint:
	golint ./...

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	$(shell rm -rf artifacts/*)
