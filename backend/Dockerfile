# Build stage
FROM golang:1.26-alpine AS builder

# Устанавливаем необходимые инструменты
RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

# Настраиваем прокси для модулей (помогает при сетевых проблемах)
ENV GOPROXY=https://proxy.golang.org,direct

# Копируем только файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости (теперь с подробным выводом)
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./main"]