FROM golang:1.23.5-alpine  

ENV WORKDIR=/app \
    GOCACHE=/tmp/.cache

WORKDIR ${WORKDIR}

COPY . .

RUN apk add --no-cache build-base

RUN go mod tidy