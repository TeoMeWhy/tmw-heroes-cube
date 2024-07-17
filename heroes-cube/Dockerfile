FROM golang:1.22 AS builder

WORKDIR /app/

COPY . .

WORKDIR /app/api

RUN go build api.go

EXPOSE 8085

CMD ["./api"]