FROM golang:1.22.5-alpine3.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o main .

CMD ["./main"]

