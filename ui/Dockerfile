FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD . /go/ui/
WORKDIR /go/ui/
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o uiApp


FROM alpine
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY  --from=builder /go/ui/uiApp /uiApp
RUN chmod +x /uiApp
EXPOSE 8084
ENV BOOKAPI_ENDPOINT=${BOOKAPI_ENDPOINT}
CMD /uiApp -apiEndpoint ${BOOKAPI_ENDPOINT}
