GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o release/linux/cadolime
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o release/win/cadolime.exe
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o release/mac/cadolime