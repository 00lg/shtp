FROM golang:1.23.2

WORKDIR /app

COPY go.mod *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /example-server

EXPOSE 8081

CMD ["/example-server"]
