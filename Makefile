bin/simple-http-server: $(wildcard *.go)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -ldflags="-s -w" -o $@ .
