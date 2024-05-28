FROM golang:latest
WORKDIR /code
COPY .  /code
CMD go run .