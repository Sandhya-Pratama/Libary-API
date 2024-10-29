# Menggunakan image Golang untuk build environment
FROM golang:1.20-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum, lalu download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh kode sumber ke dalam container
COPY . .

# Build aplikasi
RUN go build -o author-service cmd/server/main.go

# Stage kedua untuk membuat image yang lebih ringan
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary dari stage builder ke stage akhir
COPY --from=builder /app/author-service .

# Copy file .env ke dalam container
COPY .env .env

# Expose port untuk Author Service
EXPOSE 8081

# Jalankan aplikasi
CMD ["./author-service"]
