FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o author-service main.go

EXPOSE 50052

CMD ["./author-service"]