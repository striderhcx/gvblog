#!/usr/bin/env bash

set -e

# 编译mysql镜像：blog-mysql
docker build -t blog-mysql -f deploy/mysql.dockerfile .
# 编译web镜像: blog-web
docker build -t blog-web -f deploy/web.dockerfile .

# docker-compose　起容器
docker-compose up -d


