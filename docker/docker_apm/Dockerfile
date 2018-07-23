FROM nginx:alpine

LABEL maintainer "Denny<denny@dennyzhang.com>"

USER root

RUN apk add --update curl

HEALTHCHECK --interval=5m --timeout=3s \
            CMD curl -I http://localhost:80 | grep "HTTP/1.1 200 OK" || exit 1
