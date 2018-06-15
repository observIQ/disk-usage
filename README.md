# disk-usage
Alerts based on disk usage - run with cron or windows tasks

## Building
Cross compile for Windows or Linux
```
go get github.com/bluemedorapublic/gopsutil

env GOOS=windows GOARCH=amd64 go build -v -o disk-usage-win.exe
env GOOS=linux GOARCH=amd64 go build -v -o disk-usage.bin
```

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
