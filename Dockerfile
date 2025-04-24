# syntax=docker/dockerfile:1.4

# Étape de build
FROM golang:1.21 AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o placeholder-api

# Étape de prod
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/placeholder-api /app/placeholder-api

EXPOSE 8080
CMD ["/app/placeholder-api"]
