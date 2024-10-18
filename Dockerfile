FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache curl tor

COPY --from=build /app/main /app/main

COPY ./torrc /etc/tor/torrc

EXPOSE 9050 9051

CMD tor & /app/main
