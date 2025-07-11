# Dockerfile para Go 1.23 rodando em Ubuntu arm64
FROM --platform=linux/arm64 golang:1.23.8-alpine3.19

WORKDIR /app

COPY . .

RUN apk update && apk upgrade --no-cache && apk add --no-cache ca-certificates

RUN go mod download

RUN GOARCH=arm64 go build -o hello-api

CMD ["./hello-api"]
