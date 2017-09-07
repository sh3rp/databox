all: glide protobuf test serverdarwin serverlinux clientdarwin clientlinux userdbdarwin userdblinux

darwin: glide protobuf test serverdarwin clientdarwin userdbdarwin

darwinfast: protobuf serverdarwin clientdarwin userdbdarwin

linux: glide protobuf test serverlinux clientlinux userdblinux

linuxfast: protobuf serverlinux clientlinux userdblinux

test:
	go test db/* -cover -v
	go test search/* -cover -v
	go test auth/* -cover -v

serverdarwin:
	GOOS=darwin GOARCH=amd64 go build -o bawx-server.darwin cmd/server/server.go

serverlinux:
	GOOS=linux GOARCH=amd64 go build -o bawx-server.linux cmd/server/server.go

clientdarwin:
	GOOS=darwin GOARCH=amd64 go build -o bawx.darwin cmd/client/main.go

clientlinux:
	GOOS=linux GOARCH=amd64 go build -o bawx.linux cmd/client/main.go

userdbdarwin:
	GOOS=darwin GOARCH=amd64 go build -o bawx-userdb.darwin cmd/userdb/userdb.go

userdblinux:
	GOOS=linux GOARCH=amd64 go build -o bawx-userdb.linux cmd/userdb/userdb.go

protobuf:
	protoc --proto_path=msg msg/box.proto --go_out=plugins=grpc:msg

glide:
	glide update
	glide install

.PHONY: server client protobuf glide