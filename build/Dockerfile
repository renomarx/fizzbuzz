FROM golang:1.19-bullseye AS builder

RUN apt-get update
RUN apt-get install -y tzdata

ENV GO111MODULE=on

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY go.mod /go/src/app/
COPY go.sum /go/src/app/
RUN go mod download

# Because it needs to compile so it's a bit long
RUN go build "github.com/mattn/go-sqlite3"

COPY . /go/src/app/
WORKDIR /go/src/app/cmd/fizzbuzz

RUN go install

WORKDIR /go/src/app

# Final image
FROM debian:bullseye-slim

RUN apt-get update
# Generally useful
RUN apt-get install -y tzdata ca-certificates curl

# dbmate: A lightweight, framework-agnostic database migration tool
RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
RUN chmod +x /usr/local/bin/dbmate

COPY --from=builder /go/bin/fizzbuzz /usr/bin/fizzbuzz

RUN mkdir -p /app
WORKDIR /app
COPY db /app/db

RUN mkdir -p /data

# CMD fizzbuzz
CMD dbmate wait && dbmate up && fizzbuzz
