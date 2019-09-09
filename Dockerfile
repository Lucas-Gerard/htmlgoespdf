FROM golang:1.13-alpine

ADD . /home

WORKDIR /home

RUN \
   apk add --no-cache bash git openssh chromium && \
   go get -u github.com/labstack/echo/... && \
   go get -u github.com/sirupsen/logrus/... && \
   go get -u github.com/google/uuid/...

CMD go run main.go

EXPOSE 8080
