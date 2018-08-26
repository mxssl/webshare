BINARY_NAME=webshare

.PHONY: all build clean test dep build-linux build-darwin build-windows

all: build

cross-compilation: build-linux build-darwin build-windows

build:
	go build -o ${BINARY_NAME} -v

clean:
	rm -f ${BINARY_NAME}

lint:
	golangci-lint run -v

dep:
	dep ensure

# Cross compilation
build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux-amd64 -v
build-darwin:
	env GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME}-darwin-amd64 -v
build-windows:
	env GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-windows-amd64 -v
