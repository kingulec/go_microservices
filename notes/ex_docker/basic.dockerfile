FROM golang:1.21 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY /app ./app
COPY tests ./tests

RUN go build -o webhook-server ./app

EXPOSE 8443
CMD  ["go", "run", "app/main.go", "8443"]