FROM golang:1.24.1 AS builder

WORKDIR /app

COPY . ./

RUN go mod tidy

WORKDIR /app/cmd

RUN go build -o /app/avito


FROM ubuntu:latest

WORKDIR /root

COPY --from=builder /app/avito .

COPY config.example.yaml /root/

EXPOSE 8080

CMD ["/root/avito"]