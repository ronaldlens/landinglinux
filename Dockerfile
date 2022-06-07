# syntax=docker/dockerfile:1

FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY main.go ./
RUN go build -o landinglinux

FROM alpine:latest

WORKDIR /

COPY --from=build /landinglinux /landinglinux
COPY ./assets/* /assets
COPY ./data/* /data
COPY ./index.html /

WORKDIR /
EXPOSE 80
ENTRYPOINT ["/landinglinux"]