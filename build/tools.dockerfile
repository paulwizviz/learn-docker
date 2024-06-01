ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}-alpine${OS_VER} AS builder

WORKDIR /opt

COPY ./cmd/network ./cmd/network
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o ./build/network ./cmd/network

FROM alpine:${OS_VER}

COPY --from=builder /opt/build/network /usr/local/bin/network

