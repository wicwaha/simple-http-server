VERSION := $(shell git describe --always --dirty --tags)

bin/shttp: $(wildcard *.go)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -ldflags="-s -w -X main.Version=${VERSION}" -o $@ .

package: shttp.tar.xz
shttp.tar.xz: bin/shttp
	tar -cJf $@ -C bin shttp
