FROM golang:1.20.12-bullseye AS builder

# basic packages needed
RUN set -eux; \
    apt-get update && apt-get install -y git

# env for go workdir and mod vendor
ENV GOPATH /go
ENV GO_WORKDIR $GOPATH/src/bk_reptile
ENV GOFLAGS=-mod=vendor

# claim workdir and move to workdir loc
WORKDIR $GO_WORKDIR

# copy files into workdir
ADD . $GO_WORKDIR

RUN go mod vendor

# 使用 -X main.gitcommitnum 動態將 `git_commit_num` 帶入二進制參數
RUN go build -ldflags "-X 'main.gitcommitnum=`git rev-parse --short=6 HEAD`'" -o bk_reptile

FROM ubuntu:20.04
ENV GOPATH /go
ENV GO_WORKDIR $GOPATH/src/bk_reptile
WORKDIR /app

# copy binary into container
COPY --from=builder $GO_WORKDIR/bk_reptile bk_reptile
COPY ./env.yaml.temp ./env.yaml

CMD ["./bk_reptile","--config","./env.yml"]

EXPOSE 9901