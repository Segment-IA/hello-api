# Dockerfile para Go 1.23 rodando em Ubuntu arm64
FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN apk update && apk upgrade --no-cache && apk add --no-cache ca-certificates

RUN go mod download

RUN  go build -o hello-api

CMD ["./hello-api"]
