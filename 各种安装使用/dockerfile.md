

```bash
FROM node:11
MAINTAINER easydoc.net

# 复制代码
ADD . /app

# 设置容器启动后的默认运行目录
WORKDIR /app

# 运行命令，安装依赖
# RUN 命令可以有多个，但是可以用 && 连接多个命令来减少层级。
# 例如 RUN npm install && cd /app && mkdir logs
RUN npm install --registry=https://registry.npm.taobao.org

# CMD 指令只能一个，是容器启动后执行的命令，算是程序的入口。
# 如果还需要运行其他命令可以用 && 连接，也可以写成一个shell脚本去执行。
# 例如 CMD cd /app && ./start.sh
CMD node app.js
```

编译 `docker build -t test:v1` 
```txt
-t 设置镜像名字和版本号
命令参考：https://docs.docker.com/engine/reference/commandline/build/
```
运行 docker run -p 8080:8080 --name test-hello test:v1
```txt
-p 映射容器内端口到宿主机
--name 容器名字
-d 后台运行
命令参考文档：https://docs.docker.com/engine/reference/run/
```
