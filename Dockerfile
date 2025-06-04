FROM golang:latest

WORKDIR /app
COPY . .
ENV REDIS_ADDR=cache:6379
ENV REDIS_PASSWORD=vGXCa1m)vCC+

RUN go mod tidy
RUN go mod download
RUN go build -o crawler /app/cmd/main.go

CMD ["/app/crawler"]