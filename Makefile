VERSION := $(shell git rev-parse --short HEAD)
DIST_DIRS := find * -type d -exec

build:
	go build -o iSystem -ldflags "-X main.VERSION=${VERSION}" main.go

clean:
	rm -f ./iSystem
	rm -rf ./dist

bootstrap-dist:
	go get -u github.com/mitchellh/gox

build-all:
	gox -verbose \
	-ldflags "-X main.VERSION=${VERSION}" \
	-osarch="darwin/amd64 linux/amd64 linux/386" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

dist: build-all
	cd dist && \
	$(DIST_DIRS) tar -zcf isys-${VERSION}-{}.tar.gz {} \; && \
	cd ..


.PHONY: build clean bootstrap-dist build-all dist
