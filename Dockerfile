FROM golang:1.23 AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o out

FROM debian:bookworm-slim
COPY --from=build /app/out /app
CMD ["/app"]
