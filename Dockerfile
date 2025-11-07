FROM golang:1.24-alpine

WORKDIR /app

COPY . .

# download dependencies
RUN go mod download

# build binary
RUN go build -o main ./cmd/core

EXPOSE 8002

CMD ["./main"]
