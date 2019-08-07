FROM golang:1.12

WORKDIR /app

RUN mkdir ./bin

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/piqad github.com/azzz/piqad/cmd/piqad
