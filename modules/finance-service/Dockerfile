FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o finance-service ./cmd/main.go

EXPOSE 3001

CMD ["./finance-service"]
