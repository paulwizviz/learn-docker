ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}-alpine${OS_VER} AS builder

WORKDIR /opt

COPY ./cmd/socket ./cmd/socket
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o ./build/socket ./cmd/socket

FROM alpine:${OS_VER}

COPY --from=builder /opt/build/socket /usr/local/bin/socket

ENTRYPOINT [ "socket" ]
CMD ["-server=true"]