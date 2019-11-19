# disk-usage
Alerts based on disk usage - run with cron or windows tasks

## Building
A `Makefile` is provided, and relies on Docker to be available
```
make
```

If you wish to avoid Make and Docker, you can build with
Go 1.13 on your machine
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
-c string
      Pass a slack channel (default "#some_channel")
-dryrun
      Run without sending alerts
-r string
      Pass an email recipient (default "email@bdomain.com")
-s	Enable slack by passing 'true'
-smtp-port int
      Pass an smtp listening port (default 25)
-smtp-server string
      Pass an smtp server hostname (default "smtp.domain.localnet")
-t int
      Pass a threshold as an integer (default 85)
-version
      Get current version

```

Example: Use slack, but fall back to email on slack failure
```
./disk-usage.bin \
    -t 80 \
    -s -c "#my-channel" -slack-url https://hooks.slack.com/services/mycookhere \
    -r "myemail@mydomain.com" -smtp-server smtp.mydomain.com
```
