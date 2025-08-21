.PHONY: all build linux windows clean fmt test

# Default target - build both platforms
all: linux windows

# Build both Linux and Windows binaries (alias for all)
build: all

# Linux build
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpeek .

# Windows cross-compile
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o httpeek.exe .

# Format source code
fmt:
	go fmt ./...

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f httpeek httpeek.exe

