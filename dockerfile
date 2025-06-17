# Build aşaması
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o productservice ./cmd/server

# Çalışma aşaması
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/productservice .
EXPOSE 8080
CMD ["./productservice"] 