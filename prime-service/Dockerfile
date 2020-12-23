FROM golang:1.15 as builder

WORKDIR /app

COPY . .

RUN make build

EXPOSE 8000