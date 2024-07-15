.PHONY: default di clean build

default: di build;

di:
	wire ./cmd/raybot

clean:
	rm -rf ./out

build:
	go build -o out/sui-master ./cmd/raybot

install:
	go install ./cmd/raybot

static:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./out/raybot ./cmd/raybot
