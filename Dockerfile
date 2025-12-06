# --- Этап 1: Сборщик (Builder Stage) ---
FROM golang:1.24.3 AS builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app 

# Копируем файлы go.mod и go.sum и скачиваем зависимости
COPY go.mod .
COPY go.sum .
RUN go mod download 

#Копируем весь остальной код/папки
COPY . .
# Собираем исполняемый файл. CGO_ENABLED=0 для статической сборки, -o app для имени файла
RUN CGO_ENABLED=0 go build -o app ./cmd

# --- Этап 2: Финальный образ (Final Stage) ---
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/app /app/app


# Объявляем порт (информативно, не открывает его)
EXPOSE 8080
# Команда, которая запускается при старте контейнера
CMD [ "/app/app" ]
