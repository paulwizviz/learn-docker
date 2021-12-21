ARG GO_VER

FROM golang:${GO_VER}

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN env GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/restserver ./cmd/restserver

EXPOSE 9090

CMD ["/usr/local/bin/restserver"]