FROM  ubuntu:latest
MAINTAINER  yuanjun.zeng
CMD go build
RUN mkdir /more
ADD ./more-for-redis /bin/
EXPOSE "8000"
CMD  /bin/more-for-redis