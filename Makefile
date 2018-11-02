.PHONY: build test clean build_all release

build:
	go build -o dist/wait-for main.go

test:
	go test ./...

clean:
	rm dist/*

build_all:
	GOARCH=arm GOOS=linux go build -o dist/wait-for_linux_arm main.go
	GOARCH=386 GOOS=linux go build -o dist/wait-for_linux_386 main.go
	GOARCH=amd64 GOOS=linux go build -o dist/wait-for_linux_amd64 main.go
	GOOS=darwin go build -o dist/wait-for_darwin main.go
	GOARCH=386 GOOS=windows go build -o dist/wait-for_windows_386 main.go
	GOARCH=amd64 GOOS=windows go build -o dist/wait-for_windows_amd64 main.go

release: test clean build_all
