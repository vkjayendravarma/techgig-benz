FROM ubuntu:20.04
LABEL maintainer="Jayendra Varma<vkjayendravarma@gmail.com>"

RUN apt-get update

FROM golang:1.17-alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build .

EXPOSE 4000
CMD ["./techgig-benz-v2"]