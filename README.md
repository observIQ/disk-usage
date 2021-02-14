# disk-usage
Alerts based on disk usage - run with cron or windows tasks

[![Build Status](https://travis-ci.com/BlueMedoraPublic/disk-usage.svg?branch=master)](https://travis-ci.com/BlueMedoraPublic/disk-usage)
[![Go Report Card](https://goreportcard.com/badge/github.com/BlueMedoraPublic/disk-usage)](https://goreportcard.com/report/github.com/BlueMedoraPublic/disk-usage)

## Usage
Pass `--help`
```
Usage of ./disk-usage:
  -alert-type string
    	Alert type to use. Defaults to slack for backwards compatability, falls back on Stdout if slack params are not set (default "slack")
  -c string
    	Slack channel
  -dryrun
    	Run without sending alerts
  -hostname string
    	Set the hostname
  -log-level string
    	Set log level (error, warning, info, trace) (default "info")
  -slack-url string
    	Slack webhook url
  -t int
    	Disk usage percentage that should trigger an alert (default 85)
  -version
    	Print version
```

Example: Use slack:
```
./disk-usage.bin \
    -t 80 \
    -c "#my-channel" \
    -slack-url https://hooks.slack.com/services/mycookhere
```

## Building
A `Makefile` is provided, and relies on Docker to be available
```
make
```

If you wish to avoid Make and Docker, you can build with
Go 1.15 on your machine
```
go install github.com/mitchellh/gox

env CGO_ENABLED=0 \
$GOPATH/bin/gox \
    -os='!netbsd !openbsd !darwin'  \
    -output "artifacts/disk-usage-{{.OS}}-{{.Arch}}" \
    ./...
```

Both build options will output binaries in `artifacts/`
