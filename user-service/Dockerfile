FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o user-service cmd/server/main.go

CMD ["./user-service"]