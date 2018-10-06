FROM golang:latest

COPY . /app

WORKDIR /app

RUN go get ./...

RUN go run main.go
