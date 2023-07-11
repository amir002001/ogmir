# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /go/src/

COPY . .

RUN go build -o ogapp

ENTRYPOINT [ "/go/src/ogapp" ]
