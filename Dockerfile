# syntax=docker/dockerfile:1.4

# Étape de build
FROM golang:1.23 AS builder

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

RUN if [ -d "/app/placeholder-api" ]; then rm -rf /app/placeholder-api; fi && mkdir -p /app/placeholder-api/fonts
RUN mkdir -p /app/placeholder-api /app/placeholder-api/fonts
COPY fonts /app/placeholder-api/fonts


