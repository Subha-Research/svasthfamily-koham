# syntax=docker/dockerfile:1

FROM golang:1.19.1-alpine as base
# FROM golang:1.16 as base

# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /pariwar-koham
EXPOSE 8080
CMD [ "/pariwar-koham" ]