.PHONY: build
build:
	go mod vendor 
	GOPRIVATE=git.cafebazaar.ir/infrastructure CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on	go build -mod=vendor -a -o bin/order-viewer ./cmd/*.go

