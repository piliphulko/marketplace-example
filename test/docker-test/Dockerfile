# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

# installing git package without storing cache in image
RUN apk update && apk add --no-cache git

WORKDIR /app

# layer caching to improve the speed of subsequent builds
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -o server ./test/docker-test


FROM alpine
# RUN apk add shadow

# RUN useradd -m user_main

COPY --from=builder /app/server /usr/local/bin/server
COPY /config/config.yaml /usr/local/bin/config.yaml

# RUN chown user_main:user_main ./app
# USER user_main

ENTRYPOINT ["/usr/local/bin/server"]
