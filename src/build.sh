env GOOS=freebsd GOARCH=amd64 go build -v -o disk-usage-bsd.bin
env GOOS=windows GOARCH=amd64 go build -v -o disk-usage-win.exe
env GOOS=linux GOARCH=amd64 go build -v -o disk-usage.bin
