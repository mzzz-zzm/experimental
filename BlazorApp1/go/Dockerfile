FROM golang:1.20-alpine

WORKDIR /app

RUN mkdir proto

COPY go/proto/ ./proto/
COPY go/go.mod .
COPY go/go.sum .
COPY go/main.go .

RUN \
    go get -u && \
    go mod tidy

ENTRYPOINT [ "go","run","main.go" ]
