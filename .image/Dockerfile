FROM golang:1.21

RUN go install github.com/cosmtrek/air@latest

WORKDIR /usr/src/app

COPY ./app .

RUN go mod tidy