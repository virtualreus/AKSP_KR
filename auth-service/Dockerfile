# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth ./cmd/auth

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/auth /app/auth
EXPOSE 8081
ENTRYPOINT ["/app/auth"]

