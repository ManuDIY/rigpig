SHELL_IMAGE=golang:1.10.3
GIT_SHA=$(shell git rev-parse --verify HEAD | cut -c1-6)
VERSION=$(shell cat VERSION)
PROJECT="rigpig"
GITREPO="rigpig"

LDFLAGS="-X ${GITREPO}/common.Version=$(VERSION) -X ${GITREPO}/common.CommitID=$(GIT_SHA)" \

default: clean deps build

all: clean deps lint gofmt gotest build-darwin build-windows64 build-linux64 build-linux386

deps:
	go get github.com/spf13/cobra
	go get github.com/spf13/viper
	go get github.com/gizak/termui

clean:
	rm -rf bin/*


build:
	go build -o bin/${PROJECT} -ldflags $(LDFLAGS)

lint:
	$(GOPATH)/bin/golint ./...

gofmt:
	gofmt -s -w .

gotest:
	go test

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/${PROJECT} -ldflags $(LDFLAGS)

build-windows64:
	GOOS=windows GOARCH=amd64 go build -o bin/windows_x64/${PROJECT}.exe -ldflags $(LDFLAGS)

build=windows86
	GOOS=windows GOARCH=386 go build -o bin/windows_x86/${PROJECT}.exe -ldflags $(LDFLAGS)

build-linux64:
	GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/${PROJECT} -ldflags $(LDFLAGS)

build-linux386:
	GOOS=linux GOARCH=386 go build -o bin/linux_386/${PROJECT} -ldflags $(LDFLAGS)


