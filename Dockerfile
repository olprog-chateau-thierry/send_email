FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git && apk add git

WORKDIR /

RUN mkdir /app

COPY . .

RUN go build -o /app/email

ENTRYPOINT ["/app/email"]
