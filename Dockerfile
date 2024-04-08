FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./tmp/main ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/tmp/main /app/main
EXPOSE 3000

CMD ["./main"]
