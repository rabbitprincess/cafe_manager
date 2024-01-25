# syntax=docker/dockerfile:1

FROM golang:1.21 AS builder
RUN apk update && apk upgrade --no-cache && apk add make

