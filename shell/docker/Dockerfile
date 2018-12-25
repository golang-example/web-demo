FROM beyondyinjl/alpine-http
MAINTAINER jinliang <jinliang@xxxx.com>
ENV BUILDTIME 2018122511

RUN mkdir -p /data/logs/web-demo/
COPY templates /templates
ADD web-demo /opt/web-demo
ADD docker-entrypoint.sh /docker-entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["./docker-entrypoint.sh"]
