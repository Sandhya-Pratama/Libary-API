FROM golang:latest


WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o book-service cmd/server/main.go

CMD ["./book-service"]