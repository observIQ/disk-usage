VERSION := $(shell cat cmd/root.go | grep 'const version' | cut -c 25- | tr -d '"')

$(shell mkdir -p artifacts)

build: clean
	$(info building disk-usage ${VERSION})

	@docker build \
		-f docker/Dockerfile \
	    --no-cache \
	    --build-arg version=${VERSION} \
	    -t disk-usage:${VERSION} .

	@docker create -ti --name disk-usageartifacts disk-usage:${VERSION} bash && \
		docker cp disk-usageartifacts:/disk-usage/artifacts/. artifacts/

	# cleanup
	@docker rm -fv disk-usageartifacts &> /dev/null

lint:
	golint ./...

fmt:
	go fmt ./...

clean:
	$(shell rm -rf artifacts/*)
