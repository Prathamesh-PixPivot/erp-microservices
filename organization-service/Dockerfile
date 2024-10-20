FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o organization-service ./cmd/main.go

EXPOSE 3003

CMD ["./organization-service"]
