.PHONY: test coverage coverage-html build run compile

PROJECT := crud-http3
MAIN_CLIENT := cmd/client.go
MAIN_SERVER := cmd/server/server.go

test:
	go test -race -v ./...

coverage:
	go test -cover -v ./...

coverage-html:
	go test -covermode=count -coverpkg=./... -coverprofile coverage/cover.out -v ./...
	go tool cover -html coverage/cover.out -o coverage/cover.html

client:
	go build -o bin/client $(MAIN_CLIENT)

server:
	echo "generate tls certificate"
	go run cmd/server/generate_cert.go -host localhost
	go build -o bin/server $(MAIN_SERVER)

compile-client:
	echo "Compiling for every windows, linux and mac os x86_64 platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 $(MAIN_CLIENT)
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 $(MAIN_CLIENT)
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64.exe $(MAIN_CLIENT)

compile-server:
	echo "Compiling for every windows, linux and mac os x86_64 platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 $(MAIN_SERVER)
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 $(MAIN_SERVER)
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64.exe $(MAIN_SERVER)
