## gvblog
```
体验地址：

http://120.78.76.74/ihxnblog
```

### golang + vue搭建的支持markdown和代码高亮的小blog

```
1. 后台直接用golang的gin框架, 数据库orm用的是gorm; 前端的vue一把梭.

2. 这个blog的契机是618买了阿里云3年的小机器+原来用python和bootstrap写的那个太慢了，用着难受

3. 一直在学golang,做点什么出来玩玩吧.
```


### 后端运行

``` 
1. 请先确保安装了docker和docker-compose

2. ./buid_and_deploy.sh

3. docker-compose up -d

4. curl 0.0.0.0:9999/test 能正常访问，就部署好了
```


### 前端运行
```
1. 本地安装node

2. 切到前端目录npm run install

3. npm run serve
```