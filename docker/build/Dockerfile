FROM golang:1.7.1-alpine

MAINTAINER Cheng-Lung Sung <clsung@gmail.com>

RUN apk add --update git
RUN go get -u github.com/clsung/gingo

CMD ${GOPATH}/bin/gingo
