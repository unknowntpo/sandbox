# Golang Server's Dockerfile
FROM golang:1.14.6-alpine

# Because we need git while downloading golang
RUN apk add --no-cache git

# Install the module connected from golang to postgreSQL 
RUN go get github.com/lib/pq
