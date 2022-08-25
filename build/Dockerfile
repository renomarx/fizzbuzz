FROM golang:1.18-bullseye AS builder

RUN apt-get update
RUN apt-get install -y tzdata

ENV GO111MODULE=on

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY go.mod /go/src/app/
COPY go.sum /go/src/app/
RUN go mod download

COPY . /go/src/app/
WORKDIR /go/src/app/cmd/fizzbuzz

RUN go install

WORKDIR /go/src/app

# Final image
FROM debian:bullseye-slim

RUN apt-get update
# Generally useful
RUN apt-get install -y tzdata ca-certificates curl

COPY --from=builder /go/bin/fizzbuzz /usr/bin/fizzbuzz

RUN mkdir -p /app
WORKDIR /app

CMD fizzbuzz
