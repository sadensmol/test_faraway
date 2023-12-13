FROM golang:1.21-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -mod=mod -o /go/bin/test_faraway ./cmd/main.go

FROM alpine:latest

RUN rm -f /var/cache/apk/* && rm -rf /root/.cache

WORKDIR /app
EXPOSE 8080
COPY --from=builder /go/bin/test_faraway .


