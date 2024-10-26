# Gunakan Go image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Install dependensi
RUN go mod download

# Build aplikasi
RUN go build -o book-service main.go

# Expose port 8080
EXPOSE 8080

# Jalankan aplikasi
CMD ["./book-service"]
