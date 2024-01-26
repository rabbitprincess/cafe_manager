# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder
RUN apk update && apk upgrade --no-cache && apk add make
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD . .
RUN make all

FROM alpine
COPY --from=builder /app/bin/* /usr/local/bin/
CMD ["cafe_manager"]
