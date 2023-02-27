FROM golang:latest

WORKDIR /go/src/github.com/nych/docker-in-a-nutshell/demos/builder-pattern

COPY ./main.go ./

RUN go build -o main main.go
