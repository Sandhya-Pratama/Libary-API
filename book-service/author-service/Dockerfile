FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o author-service cmd/server/main.go

CMD ["./author-service"]