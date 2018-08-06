FROM golang:1.10-alpine3.8 AS compiler

COPY . /go/src/github.com/IndraGunawan/gosample

RUN apk --no-cache add curl \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN cd /go/src/github.com/IndraGunawan/gosample \
    && dep ensure \
    && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build app/web/main.go

FROM alpine:3.8

RUN mkdir /gosample
COPY --from=compiler /go/src/github.com/IndraGunawan/gosample/main /gosample

EXPOSE 8080
CMD ["/gosample/main"]

