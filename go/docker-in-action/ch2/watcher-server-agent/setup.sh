#!/bin/bash

# Remove old container
docker rm -f web mailer agent

# Run the server
docker run --rm --detach \
    --name web nginx:latest

# Run the mailer
docker run --rm -d \
    --name mailer \
    dockerinaction/ch2_mailer

# Run the agent
# https://github.com/dockerinaction/ch2_agent
docker run -it \
    --name agent \
    --link web:insideweb \
    --link mailer:insidemailer \
    dockerinaction/ch2_agent
