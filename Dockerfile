FROM  ubuntu:latest
MAINTAINER  yuanjun.zeng
CMD go build
RUN mkdir /more
ADD ./moreRedis /bin/
EXPOSE "8000"
CMD  /bin/moreRedis