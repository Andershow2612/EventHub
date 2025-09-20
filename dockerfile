# Dockerfile
FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/api ./cmd

FROM debian:bullseye-slim
WORKDIR /app

COPY --from=builder /app/api .
COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

EXPOSE 8080
USER 0
# Espera o DB antes de iniciar a API
CMD ["/app/wait-for-it.sh", "db:5432", "--", "./api"]
