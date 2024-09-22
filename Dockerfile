# Используем базовый образ с Go для сборки
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем все файлы проекта
COPY . .

# Собираем бинарный файл Go приложения (указываем директорию cmd/interactions)
RUN go build -o main ./cmd/interactions

# Финальный этап: минимальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию в финальном контейнере
WORKDIR /app

# Устанавливаем нужные зависимости (например, libc6-compat для исполнения бинарников Go)
RUN apk --no-cache add libc6-compat

# Копируем бинарный файл из стадии сборки
COPY --from=builder /app/main .

# Проверяем содержимое директории (для отладки)
RUN ls -la

# Устанавливаем права на выполнение
RUN chmod +x ./main

# Указываем порт, который будет использовать приложение
EXPOSE 8088

# Указываем команду для запуска приложения
CMD ["./main"]
