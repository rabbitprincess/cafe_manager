# syntax=docker/dockerfile:1

FROM golang:1.21 AS builder
RUN apt-get update && apt-get install -y make
WORKDIR /app
COPY . .
RUN go build -o cafe_manager cafe_manager.go

FROM alpine
COPY --from=builder /app/cafe_manager /usr/local/bin/
CMD ["cafe_manager"]
