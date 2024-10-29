FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o borrowing-service cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/borrowing-service .

EXPOSE 8085

CMD ["./borrowing-service"]
