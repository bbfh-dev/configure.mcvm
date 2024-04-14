dev:
	go build -o dist/mcvmconf main.go

prod:
	GOOS="linux" GOARCH="amd64" go build -o dist/mcvmconf-linux-amd64 -ldflags "-s -w" main.go
	GOOS="windows" GOARCH="amd64" go build -o dist/mcvmconf-windows-amd64.exe -ldflags "-s -w" main.go
	GOOS="darwin" GOARCH="amd64" go build -o dist/mcvmconf-darwin-amd64 -ldflags "-s -w" main.go

watch:
	reflex -s -r '\.go$$' -- make dev --always-make

debug:
	dlv debug --headless --api-version=2 --listen=127.0.0.1:43000 main.go 
