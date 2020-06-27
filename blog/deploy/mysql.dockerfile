# image
FROM mysql:5.7

MAINTAINER ihxn <1225427292@qq.com>

ENV TZ=Asia/Shanghai
ENV LANG C.UTF-8

#定义会被容器自动执行的目录
ENV AUTO_RUN_DIR /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD 666

#把要执行的sql文件放到/docker-entrypoint-initdb.d/目录下，容器会自动执行这个sql
COPY ./deploy/dbinit/blog.sql $AUTO_RUN_DIR/
