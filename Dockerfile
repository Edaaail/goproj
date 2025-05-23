FROM golang:1.23.5

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o shorten-service .

EXPOSE 8080

CMD ["./shorten-service"]
