# Build stage
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main ./cmd

# Run stage
FROM debian:bullseye-slim
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
