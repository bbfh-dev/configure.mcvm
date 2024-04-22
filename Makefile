VERSION := v0.0.1
BINARY_NAME := mcvmconf

dev:
	go build -o dist/$(BINARY_NAME) main.go

prod:
	# Compile for Linux
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-linux-x64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

	# Compile for Linux (ARM v8)
	GOOS=linux GOARCH=arm64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-linux-arm64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

	# Compile for Windows
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-windows-x64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

	# Compile for Darwin (macOS)
	GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-darwin-x64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

	# Compile for Darwin (macOS, apple silicon)
	GOOS=darwin GOARCH=arm64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-darwin-arm64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

	# Compile for FreeBSD
	GOOS=freebsd GOARCH=amd64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-freebsd-x64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

	# Compile for FreeBSD (ARM v8)
	GOOS=freebsd GOARCH=arm64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-freebsd-arm64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

	# Compile for OpenBSD
	GOOS=freebsd GOARCH=amd64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-openbsd-x64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

	# Compile for OpenBSD
	GOOS=freebsd GOARCH=amd64 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" main.go
	tar -czvf dist/$(BINARY_NAME)-$(VERSION)-openbsd-x64.tar.gz -C dist $(BINARY_NAME)
	rm dist/$(BINARY_NAME)

watch:
	reflex -s -r '\.go$$' -- make dev --always-make

debug:
	dlv debug --headless --api-version=2 --listen=127.0.0.1:43000 main.go 
