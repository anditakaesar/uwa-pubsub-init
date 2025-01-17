FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /uwa-pubsub-init

FROM alpine:3.21

WORKDIR /root/

COPY --from=builder /uwa-pubsub-init .

CMD ["./uwa-pubsub-init"]