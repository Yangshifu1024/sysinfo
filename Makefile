VERSION := $(shell git rev-parse --short HEAD)-$(shell date '+%Y%m%d')
DIST_DIRS := find * -type d -exec

build:
	go build -o sysinfo -ldflags "-X main.VERSION=${VERSION}" main.go

clean:
	rm -f ./sysinfo
	rm -rf ./dist

bootstrap-dist:
	go get -u github.com/mitchellh/gox

build-all: clean
	gox -verbose \
	-ldflags "-w -s -X main.VERSION=${VERSION}" \
	-osarch="darwin/amd64 linux/amd64 linux/386" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

dist: build-all
	cd dist && \
	$(DIST_DIRS) tar -zcf sysinfo-${VERSION}-{}.tar.gz {} \; && \
	cd ..


.PHONY: build clean bootstrap-dist build-all dist
