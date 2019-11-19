FROM golang:1.13

ARG version

ADD . /disk-usage/
WORKDIR /disk-usage

RUN go test ./...

RUN go get github.com/mitchellh/gox

# build for Windows, Linux, FreeBSD
RUN \
    env CGO_ENABLED=0 \
    $GOPATH/bin/gox \
        -arch=amd64 \
        -os='!netbsd !openbsd !darwin'  \
        -output "artifacts/disk-usage-{{.OS}}-{{.Arch}}" \
        ./...

WORKDIR /disk-usage/artifacts
RUN ls | xargs -n1 sha256sum >> SHA256SUMS
