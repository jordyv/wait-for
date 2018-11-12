.PHONY: build test clean build_all release deps

build: deps
	go build -o dist/wait-for main.go

test:
	go test ./...

clean:
	rm -f dist/*

build_all: deps
	GOARCH=arm GOOS=linux go build -o dist/wait-for_linux_arm main.go
	GOARCH=386 GOOS=linux go build -o dist/wait-for_linux_386 main.go
	GOARCH=amd64 GOOS=linux go build -o dist/wait-for_linux_amd64 main.go
	GOOS=darwin go build -o dist/wait-for_darwin main.go
	GOARCH=386 GOOS=windows go build -o dist/wait-for_windows_386 main.go
	GOARCH=amd64 GOOS=windows go build -o dist/wait-for_windows_amd64 main.go

deps:
	dep ensure -v

release: test clean build_all
