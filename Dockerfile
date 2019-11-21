FROM golang:1.13

ARG version

ADD . /disk-usage/
WORKDIR /disk-usage

RUN go get github.com/mitchellh/gox

# build for Windows, Linux, FreeBSD
RUN \
    $GOPATH/bin/gox \
        -arch=amd64 \
        -os='!netbsd !openbsd !darwin'  \
        -output "artifacts/disk-usage-{{.OS}}-{{.Arch}}" \
        -ldflags '-w -extldflags "-static"' \
        ./...

WORKDIR /disk-usage/artifacts
RUN ls | xargs -n1 sha256sum >> SHA256SUMS
