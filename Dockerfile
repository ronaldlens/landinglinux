# syntax=docker/dockerfile:1

FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o landinglinux

FROM alpine:latest

WORKDIR /

COPY --from=build /app/landinglinux /landinglinux
COPY ./assets/* /assets/
COPY ./data/* /data/
COPY ./index.html /

EXPOSE 80
ENTRYPOINT ["/landinglinux"]
