FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o short-url-service ./cmd/app/main.go

FROM golang:1.21
WORKDIR /app
COPY --from=builder /app/short-url-service .

CMD ["./short-url-service"]
