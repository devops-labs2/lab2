test:
	go test ./...

build:
	GOOS=linux GOARCH=arm64 go build -trimpath -ldflags='-s -w' -o bootstrap ./cmd
	rm -f lambda-handler.zip
	zip -j lambda-handler.zip bootstrap
