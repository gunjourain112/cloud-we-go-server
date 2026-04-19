FROM golang:1.26-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -tags=sonic -o main .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY .env.example .env

EXPOSE 8081

CMD ["./main"]
