# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o source-service .

# Run stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/source-service .

EXPOSE 8080

CMD ["./source-service"]
