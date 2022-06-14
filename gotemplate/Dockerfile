FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
WORKDIR /go/src/github.com/roost-io/roost-example-latest/googlebookapi/gotemplate
RUN go build -o gotemplate

EXPOSE 5200
USER nobody
ENTRYPOINT ["/app/gotemplate"]
