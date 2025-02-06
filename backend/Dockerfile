# Этап сборки
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend

FROM golang:1.21

WORKDIR /app

COPY --from=builder /app/backend .

EXPOSE 8080

CMD ["./backend"]
