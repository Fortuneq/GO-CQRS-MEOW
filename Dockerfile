FROM golang:1.13-alpine3.11 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/tinrab/meower

COPY go.mod go.sum ./
COPY util util
COPY event event
COPY db db
COPY search search
COPY schema schema
COPY meow-service meow-service
COPY query-service query-service
COPY pusher-service pusher-service

RUN GO111MODULE=on go install github.com/elastic/go-elasticsearch/v7

# github.com/gorilla/mux v1.7.4
RUN GO111MODULE=on go install  github.com/gorilla/mux
# github.com/gorilla/websocket v1.4.1
RUN GO111MODULE=on go install github.com/gorilla/websocket
# github.com/kelseyhightower/envconfig v1.4.0
RUN GO111MODULE=on go install github.com/kelseyhightower/envconfig
# github.com/lib/pq v1.3.0
RUN GO111MODULE=on go install github.com/lib/pq

# github.com/nats-io/nats.go v1.9.1
RUN GO111MODULE=on go install github.com/nats-io/nats.go

# github.com/segmentio/ksuid v1.0.2
RUN GO111MODULE=on go install github.com/segmentio/ksuid
# github.com/tinrab/retry v1.0.0
RUN GO111MODULE=on go install github.com/tinrab/retry
# golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .