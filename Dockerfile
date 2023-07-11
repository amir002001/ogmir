# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /go/src/

RUN apk add \--update chromium && \  
    apk info \--purge 

COPY . .

RUN go build -o ogapp

ENTRYPOINT [ "/go/src/ogapp" ]
