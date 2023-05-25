### builder image ###
FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye as builder

WORKDIR /grpc-demo
COPY . .

# install deps
RUN scripts/install-taskfile.sh
RUN scripts/install-bufbuild.sh

# link statically
ENV CGO_ENABLED=0
RUN task build



### dev base image for debugging ###
FROM ubuntu:latest as dev-base

RUN apt-get update && apt-get install -y \
  iputils-ping \
  netcat \
  curl \
  && rm -rf /var/lib/apt/lists/*

# for buf curl
RUN scripts/install-bufbuild.sh



### server image ###
FROM alpine:latest as server

COPY --from=builder  /grpc-demo/cmd/server/main .

ENTRYPOINT ["/main"]

CMD ["--port", "8081"]



FROM dev-base as dev-server

COPY --from=builder  /grpc-demo/cmd/server/main .

CMD ["/main", "--port", "8081"]



### reverseproxy image ###
FROM alpine:latest as reverseproxy

COPY --from=builder  /grpc-demo/cmd/reverseproxy/main .

ENTRYPOINT ["/main"]

CMD ["--grpcServer", "localhost:8081", "--port", "8081"]



FROM dev-base as dev-reverseproxy

COPY --from=builder  /grpc-demo/cmd/reverseproxy/main .

CMD ["/main", "--grpcServer", "localhost:8081", "--port", "8081"]