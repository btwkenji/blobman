FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/blobman
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/blobman /go/src/blobman

FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/blobman /usr/local/bin/blobman
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["blobman"]
