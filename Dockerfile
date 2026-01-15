FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY /app ./app
COPY tests ./tests

RUN go build -o webhook-server ./app

FROM ubuntu:22.04 AS final
WORKDIR /app
COPY --from=builder /app/webhook-server ./webhook-server
EXPOSE 8443
ENV TLS_CERT_PATH=/certs/tls.crt
ENV TLS_KEY_PATH=/certs/tls.key
CMD  ["./webhook-server", "8443"] 