FROM alpine AS www
WORKDIR /www
ADD https://github.com/glarfs/nginx-error-pages/tarball/master /
RUN tar --strip-components=1 -xf /master -C /www/

FROM golang:1.12.1-alpine3.9 as builder

WORKDIR /go/src/github.com/nuvo/default-backend
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o main github.com/nuvo/default-backend

FROM scratch

COPY --from=builder /go/src/github.com/nuvo/default-backend/main /
COPY --from=www /www/ /assets/


CMD ["/main"]
