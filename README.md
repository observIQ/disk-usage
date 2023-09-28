# disk-usage
Alerts based on disk usage - run with cron or windows tasks

[![Build Status](https://travis-ci.com/observIQ/disk-usage.svg?branch=master)](https://travis-ci.com/observIQ/disk-usage)
[![Go Report Card](https://goreportcard.com/badge/github.com/observIQ/disk-usage)](https://goreportcard.com/report/github.com/observIQ/disk-usage)

## Usage
Pass `--help`
```
Usage of ./disk-usage:
  -alert-type string
    	Alert type to use. Defaults to slack for backwards compatibility, falls back on Stdout if slack params are not set (default "slack")
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

Example message:
```
{
    "host": {
        "name": "test",
        "address": "10.99.1.10",
        "devices": [
            {
                "name": "/dev/dm-1",
                "mountpoint": "/",
                "type": "ext4",
                "usage_percent": 87
            },
            {
                "name": "/dev/sdc2",
                "mountpoint": "/boot",
                "type": "ext4",
                "usage_percent": 29
            },
            {
                "name": "/dev/dm-3",
                "mountpoint": "/home",
                "type": "ext4",
                "usage_percent": 63
            },
            {
                "name": "/dev/sda1",
                "mountpoint": "/mnt",
                "type": "ext4",
                "usage_percent": 21
            }
        ]
    },
    "message": "devices have high usage: [/dev/dm-1 /dev/sdc2 /dev/dm-3 /dev/sda1]",
    "severity": "fatal"
}
```

## Building
A `Makefile` is provided, and relies on Docker to be available
```
make
```

If you wish to avoid Make and Docker, you can build with
Go 1.16 on your machine
```
go install github.com/mitchellh/gox

env CGO_ENABLED=0 \
$GOPATH/bin/gox \
    -os='!netbsd !openbsd !darwin'  \
    -output "artifacts/disk-usage-{{.OS}}-{{.Arch}}" \
    ./...
```

Both build options will output binaries in `artifacts/`

# Community

disk-usage is an open source project. If you'd like to contribute, take a look at our [contribution guidelines](/docs/CONTRIBUTING.md). We look forward to building with you.

## Code of Conduct

disk-usage follows the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Please report violations of the Code of Conduct to any or all [maintainers](/docs/MAINTAINERS.md).

# Other questions?

Send us an [email](mailto:support@observiq.com), or open an issue with your question. We'd love to hear from you!
