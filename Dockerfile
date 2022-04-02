FROM golang:1.16.3-alpine
WORKDIR /usr/src/go

RUN apk add --update \
    git \
    protoc

RUN go install github.com/golang/protobuf/protoc-gen-go@latest

CMD sh -c "go mod tidy && go run main.go"