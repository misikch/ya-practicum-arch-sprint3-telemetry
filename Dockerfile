# Dockerfile
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка сервисов
RUN go build -o /service ./cmd/service/main.go
RUN go build -o /worker ./cmd/worker/main.go

# Используем более легкое изображение в качестве окончательного
FROM alpine:latest

# Создаем рабочую директорию
WORKDIR /app

# Копируем собранные бинарные файлы из предыдущего этапа
COPY --from=builder /service .
COPY --from=builder /worker .
