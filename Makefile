.PHONY: all build linux windows darwin clean fmt test

# Default target - build all platforms
all: linux windows darwin

# Build all binaries (alias for all)
build: all

# Linux build
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpeek .

# Windows cross-compile
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o httpeek.exe .

# macOS build (Apple Silicon)
darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o httpeek-darwin .

# Format source code
fmt:
	go fmt ./...

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f httpeek httpeek.exe httpeek-darwin

