FROM alpine:3.6
MAINTAINER Lucas <6congyao@gmail.com>

RUN apk add -U --no-cache ca-certificates

ADD iamprototype /bin/
EXPOSE 8080

HEALTHCHECK CMD ["/bin/iamprototype", "ping"]

ENTRYPOINT ["/bin/iamprototype"]