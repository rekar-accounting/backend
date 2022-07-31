# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18.4-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

USER nonroot:nonroot

ENTRYPOINT ["/main"]