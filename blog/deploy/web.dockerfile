# image
FROM golang:latest

MAINTAINER ihxn <1225427292@qq.com>

ENV TZ=Asia/Shanghai
ENV LANG C.UTF-8

RUN mkdir -p /data/www/blog 
RUN touch /data/www/blog/app.log
COPY ./main /data/www/blog/
COPY ./config/*.yaml /data/www/blog/config/


WORKDIR /data/www/blog

ENTRYPOINT ["./main"]
