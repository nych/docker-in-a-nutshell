FROM golang:latest as builder

WORKDIR /go/src/github.com/nych/docker-in-a-nutshell/demos/builder-pattern

COPY ./main.go ./

RUN go build -o main main.go



FROM alpine:latest as multi-stage-demo

COPY --from=builder /go/src/github.com/nych/docker-in-a-nutshell/demos/builder-pattern/main ./

CMD ["./main"]