FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o graphql-gateway ./cmd/main.go

EXPOSE 3002

CMD ["./graphql-gateway"]
