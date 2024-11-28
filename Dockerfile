FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/web

FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=builder /app/main .

COPY ./ui ./ui

EXPOSE 4000

CMD ["./main", "-addr", ":4000", "-dsn", "web:7777@tcp(mysql:3306)/snippetbox?parseTime=true"]
