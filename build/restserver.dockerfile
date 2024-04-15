FROM golang:1.20.6-alpine3.18 AS builder

WORKDIR /opt

COPY ./cmd/restserver ./cmd/restserver
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o ./bin/restserver ./cmd/restserver/main.go

FROM alpine:3.18.2

COPY --from=builder /opt/bin/restserver /usr/local/bin/restserver

CMD /usr/local/bin/restserver