FROM golang:1.24.1-alpine3.21 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/main ./cmd/main.go

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/main .

ENV GODEBUG="http2debug=1"
EXPOSE 9000
CMD ["./main", "-server-port=9000"]
