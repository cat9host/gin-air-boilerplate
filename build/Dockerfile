ARG TARGET_DIR=/app
ARG GOBIN=/.bin

FROM golang:1.22.1-alpine

ARG TARGET_DIR
ARG GOBIN

RUN apk add --update make git musl-dev gcc libc-dev binutils-gold pkgconfig czmq-dev

ENV GO111MODULE=on
ENV GOPATH=/go
WORKDIR /go/src/github.com/cat9host/gin-air-boilerplate

COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
ADD cmd cmd
ADD internal internal
RUN go mod download
RUN GOBIN=$GOBIN make install-tools
RUN TARGET_DIR=$TARGET_DIR make build

ENTRYPOINT make watch
