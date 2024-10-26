# Gunakan Golang sebagai base image
FROM golang:1.19-alpine

# Setel direktori kerja
WORKDIR /app

# Copy go mod dan sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh kode sumber
