FROM golang:1.9-alpine

WORKDIR /go/src/apiServer
ADD . .
RUN go install

ENTRYPOINT /go/bin/apiServer
EXPOSE 8080
