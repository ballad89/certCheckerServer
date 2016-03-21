FROM golang:1.6

MAINTAINER husayn arrah

ADD . /go/src/certcheckerServer/

WORKDIR /go/src/certcheckerServer

RUN chmod a+x /go/src/certcheckerServer/*.sh && mkdir -p /data/db

RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10 && \
    echo "deb http://repo.mongodb.com/apt/debian jessie/mongodb-enterprise/3.3 main" >> /etc/apt/sources.list && \
    apt-get update && \
    apt-get install mongodb-enterprise-unstable-server redis-server -y --force-yes && \
    apt-get clean

RUN go install

EXPOSE 3000

VOLUME /data/db

CMD /go/src/certcheckerServer/startup.sh