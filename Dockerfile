FROM golang:alpine AS builder

RUN mkdir /app

COPY . /app

RUN apk update && apk add --no-cache git && apk add git

WORKDIR /app

RUN git clone https://github.com/olprog-chateau-thierry/send_email.git

RUN go build -o /app/send_email

FROM scratch

COPY --from=builder /app/send_email /app/send_email

EXPOSE 8080

ENTRYPOINT ["/app/send_email"]
