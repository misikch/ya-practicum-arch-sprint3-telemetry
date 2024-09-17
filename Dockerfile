# Dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main cmd/service/main.go

CMD ["./main"]
