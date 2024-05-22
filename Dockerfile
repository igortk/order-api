FROM golang:1.21-alpine AS builder

WORKDIR /order-api

COPY ./ ./

RUN go build -o main .

RUN go mod download

EXPOSE 8080

CMD ["./order-api"]