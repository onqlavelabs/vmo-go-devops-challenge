FROM golang:1.18 as build

ARG CGO_ENABLE=0
ARG GO111MODULE=on

RUN mkdir -p /go/src/github.com/dinhtp/vmo-go-devops-challenge/application
WORKDIR ~/go/src/github.com/dinhtp/vmo-go-devops-challenge/application
COPY . .

RUN GOOS=linux go build -o /go/bin/programming-challenge .

FROM ubuntu:22.04

EXPOSE 9090

WORKDIR /app

COPY ./db.json db.json
COPY ./docs docs/
COPY --from=build /go/bin/programming-challenge /usr/bin/

CMD ["programming-challenge", "serve"]