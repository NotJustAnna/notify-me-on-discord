# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum  *.go ./

RUN go mod download && go build -o /notify-me-on-discord

EXPOSE 8080

CMD [ "/notify-me-on-discord" ]