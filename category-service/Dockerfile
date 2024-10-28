FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o category-service cmd/server/main.go

CMD ["./category-service"]