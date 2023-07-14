# syntax=docker/dockerfile:1

FROM golang:alpine

ENV CGO_ENABLED=1

WORKDIR /go/src/

RUN apk add \--update chromium && \  
    apk add build-base && \
    apk info \--purge 

COPY . .

RUN go build -o ogapp

ENTRYPOINT [ "/go/src/ogapp" ]
