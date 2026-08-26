FROM golang:1.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY cmd ./cmd

RUN go build -o app ./cmd

EXPOSE 8080

CMD ["./app"]