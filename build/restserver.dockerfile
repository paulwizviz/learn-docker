ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}-alpine${OS_VER} AS builder

WORKDIR /opt

COPY ./cmd/restserver ./cmd/restserver
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o ./bin/restserver ./cmd/restserver/main.go

FROM alpine:${OS_VER}

RUN apk update && \
    apk upgrade && \
    apk --no-cache add curl

COPY --from=builder /opt/bin/restserver /usr/local/bin/restserver

CMD /usr/local/bin/restserver