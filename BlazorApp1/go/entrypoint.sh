#!/bin/sh
protoc -Iprotobuf --go_opt=paths=source_relative --go_out=./proto --go-grpc_opt=paths=source_relative --go-grpc_out=./proto ./protobuf/**/*.proto
go get -u
go mod tidy
