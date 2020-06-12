ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir /api
WORKDIR /api

COPY . .
RUN go mod download

RUN go build -o ./app ./main.go

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir /api
WORKDIR /api
COPY --from=builder /api/. .

EXPOSE 8080

ENTRYPOINT ["./app"]