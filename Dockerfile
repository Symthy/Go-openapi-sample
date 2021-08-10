FROM golang:latest

RUN mkdir /go/src/work
RUN apt-get update && apt-get install
RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

WORKDIR /go/src/work

CMD ["/bin/bash"]
