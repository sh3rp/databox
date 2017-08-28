all: glide protobuf test server client

test:
	go test db/* -cover

server:
	GOOS=darwin GOARCH=amd64 go build -o bawx-server.darwin cmd/server/server.go
	GOOS=linux GOARCH=amd64 go build -o bawx-server.linux cmd/server/server.go

client:
	GOOS=darwin GOARCH=amd64 go build -o bawx.darwin cmd/client/main.go
	GOOS=linux GOARCH=amd64 go build -o bawx.linux cmd/client/main.go

protobuf:
	protoc --proto_path=msg msg/box.proto --go_out=plugins=grpc:msg

glide:
	glide update
	glide install

.PHONY: server client protobuf glide