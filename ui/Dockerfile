FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD . /go/ui/
WORKDIR /go/ui/
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ui-app


FROM alpine
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY  --from=builder /go/ui/ui-app /ui-app
EXPOSE 8084
USER nobody
ENTRYPOINT ["/ui-app"]
