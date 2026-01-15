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
CMD  ["./webhook-server", "8443"] 