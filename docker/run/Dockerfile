FROM alpine:3.4

MAINTAINER Cheng-Lung Sung <clsung@gmail.com>

ADD gingo /gingo

# in case we need to verify log
VOLUME /var/log
EXPOSE 8080

CMD ["/gingo"]
