# Builder
FROM golang:1.23-alpine3.21 as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY . .
COPY .env /app/.env

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /app/engine /app/
COPY .env /app/.env

CMD /app/engine