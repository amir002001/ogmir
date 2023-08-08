# syntax=docker/dockerfile:1

FROM golang:alpine as build

ENV CGO_ENABLED=1

WORKDIR /go/src/

RUN apk add \--update build-base

COPY . .

RUN go build -o ogapp



FROM alpine:latest
WORKDIR /srv

COPY --from=build /go/src/ogapp ./ogapp
COPY ./www ./www
COPY ./index.template.html ./

RUN apk add --update chromium && apk info --purge

ENTRYPOINT [ "/srv/ogapp" ]
