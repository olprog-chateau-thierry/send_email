FROM --platform=linux/amd64 golang:alpine AS builder

RUN apk update && apk add --no-cache git && apk add git


RUN mkdir /app

WORKDIR /app

COPY . .

RUN ls -alh

ENV CGO_ENABLED 0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -ldflags="-w -s" -o /app/email

FROM --platform=linux/amd64 scratch

WORKDIR /app

COPY --from=builder /app/email /app/email

EXPOSE 8080

ENTRYPOINT ["./email"]
