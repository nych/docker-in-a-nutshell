# base
FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye as base

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d



# devcontainer image
FROM base as devcontainer

USER vscode



# builder image
FROM base as builder

WORKDIR /go/src/github.com/nych/docker-in-a-nutshell/demos/extended-devcontainer

COPY . .

RUN task build



# greeter image
FROM alpine:latest as greeter

COPY --from=builder /go/src/github.com/nych/docker-in-a-nutshell/demos/extended-devcontainer/build/main .

ENTRYPOINT ["/main"]

CMD ["--name", "Paul Atreides"]