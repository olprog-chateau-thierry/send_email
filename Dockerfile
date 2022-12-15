FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git && apk add git

WORKDIR /

RUN mkdir /app

RUN git clone https://github.com/olprog-chateau-thierry/send_email.git

WORKDIR send_email

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/send_email

RUN ls -alh /app

FROM scratch

COPY --from=builder /app/send_email /email

EXPOSE 8080

ENTRYPOINT ["/email"]
