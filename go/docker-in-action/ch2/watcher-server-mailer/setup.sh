#!/bin/bash

# Remove old container
docker rm -f web mailer web_test

# Run the server
docker run --rm --detach \
    --name web nginx:latest

# Run the mailer
docker run --rm -d \
    --name mailer \
    dockerinaction/ch2_mailer

# Run the watcher
docker run --rm --interactive --tty \
    --link web:web \
    --name web_test \
    busybox:1.29 /bin/sh
