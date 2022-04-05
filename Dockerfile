FROM golang:1.16.3-alpine
WORKDIR /usr/src/go

RUN apk add --no-cache --update-cache \
    git \
    protoc

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .