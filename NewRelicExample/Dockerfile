# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -buildvcs=false -o /jarvis

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /new-relic-example /new-relic-example

USER nonroot:nonroot

EXPOSE 80

ENTRYPOINT ["/new-relic-example"]