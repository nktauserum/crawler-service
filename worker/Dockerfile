FROM golang:latest

WORKDIR /app
COPY . .
COPY go.mod .
COPY go.sum .

ENV GRPC_PORT=crawler:50000

RUN go mod tidy
RUN go mod download

RUN go build -o worker /app/worker.go

CMD ["./worker"]