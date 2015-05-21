FROM golang:1.4.2-cross
MAINTAINER peter.edge@gmail.com

RUN mkdir -p /go/src/github.com/peter-edge/go-yaml2json
ADD . /go/src/github.com/peter-edge/go-yaml2json/
WORKDIR /go/src/github.com/peter-edge/go-yaml2json
