# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o medical ./cmd/medical

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/medical /app/medical
EXPOSE 8082
ENTRYPOINT ["/app/medical"]

