FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o leads-service ./cmd/main.go

EXPOSE 3001

CMD ["./leads-service"]
