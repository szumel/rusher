VERSION?=$(shell git describe --tags)
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
LDFLAGS = -ldflags "-X github.com/szumel/rusher/internal/platform/version.VERSION=${VERSION}"

all: tests build readme

tests:
	go test -v ./...

build:
	go build ${LDFLAGS} cmd/cli/rusher.go
readme:	
	rm -f README.md
	cp README.md.template README.md
	./rusher listSteps >> README.md
