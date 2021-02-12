# disk-usage
Alerts based on disk usage - run with cron or windows tasks

[![Build Status](https://travis-ci.com/BlueMedoraPublic/disk-usage.svg?branch=master)](https://travis-ci.com/BlueMedoraPublic/disk-usage)

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
    -arch=amd64 \
    -os='!netbsd !openbsd !darwin'  \
    -output "artifacts/disk-usage-{{.OS}}-{{.Arch}}" \
    ./...
```

Both build options will output binaries in `artifacts/`

## Usage
Pass `--help`
```
Usage of ./disk-usage:
  -alert-type string
    	Alert type to use. Defaults to slack for backwards compatability, falls back on Stdout if slack params are not set (default "slack")
  -alsologtostderr
    	log to standard error as well as files
  -c string
    	Slack channel
  -dryrun
    	Run without sending alerts
  -log_backtrace_at value
    	when logging hits line file:N, emit a stack trace
  -log_dir string
    	If non-empty, write log files in this directory
  -logtostderr
    	log to standard error instead of files
  -slack-url string
    	Slack webhook urlL
  -stderrthreshold value
    	logs at or above this threshold go to stderr
  -t int
    	Disk usage percentage that should trigger an alert (default 85)
  -v value
    	log level for V logs
  -version
    	Print version
  -vmodule value
    	comma-separated list of pattern=N settings for file-filtered logging
```

Example: Use slack:
```
./disk-usage.bin \
    -t 80 \
    -c "#my-channel" \
    -slack-url https://hooks.slack.com/services/mycookhere
```
