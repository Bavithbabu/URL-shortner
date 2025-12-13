FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/main .
COPY --from=builder /build/.env .

RUN chmod +x ./main

EXPOSE 3000

CMD ["./main"]
