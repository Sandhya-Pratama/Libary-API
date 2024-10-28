FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o borrowing-service cmd/server/main.go

CMD ["./borrowing-service"]