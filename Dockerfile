FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o kbot main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/kbot .
CMD ["./kbot"]
