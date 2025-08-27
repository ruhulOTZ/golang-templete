# Stage 1: Build with Go 1.25
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o myapp ./main.go  # Adjust path to your main file.

# Stage 2: Create a slim image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 4000
CMD ["./myapp"]