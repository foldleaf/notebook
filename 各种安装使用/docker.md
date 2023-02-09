下载docker桌面版
https://www.docker.com/products/docker-desktop/

控制面板->程序->启用或关闭 windows 功能，开启 `Windows 虚拟机监控程序平台`和`适用于Linux的Windows子系统`
设置开机启动 Hypervisor
```bash
bcdedit /set hypervisorlaunchtype auto
```

确认任务管理器-性能-CPU面板-虚拟化为已开启，确保 BIOS 已开启虚拟化
使用镜像源
```bash
"registry-mirrors": ["https://registry.docker-cn.com"]
```
重启，按提示操作

示例：redis
```bash
# docker pull redis 或者 docker pull redis:5.0.5 或者
docker run -d -p 6379:6379 --name redis redis:latest
```
查看拉取成功的镜像
```bash
docker images
```
```bash
docker run --name redis -p 6379:6379 --restart=always -v $PWD/data:/data  -d redis:5.0.5 redis-server --appendonly yes daemonize yes
# 参数说明：
# 本地运行
-d
# 本地端口:Docker端口
6379:6379
# 指定驱动盘
-v
# Redis的持久化文件存储
$PWD/data
# docker的镜像名
redis
# redis服务器
redis-server
# 开启持久化
--appendonly yes
# 这个运行的镜像的名称
--name
# 守护进程
daemonize yes
#Docker启动容器就启动
--restart=always
```

启动镜像
```bash
docker start redis
```

停止正在运行的镜像(
```bash
# redis为--name的名称
docker stop redis
```
删除镜像
```bash
docker rmi <image id>
# 删除全部
docker rmi $(docker images -q)
```

获取 container 的ID或name
```bash
docker container ls -a
```

根据container的ID 停止container
```bash
docker  container  stop   CONTAINER_ID
```

根据ID或name删除container（先停止才能删除）
```bash
docker   container  rm  CONTAINER_ID  或者 docker  container  rm  CONTAINER_NAME 
```

批量操作
```bash
# 批量获取容器ID
docker container ls -a -q
# 批量获取镜像ID
docker image ls -a -q  
# 批量停止容器
docker container   stop   $(docker  container  ls   -a  -q)
# 批量删除容器
docker   container   rm  $(docker  container  ls   -a  -q)
```


```bash
# 访问容器
docker exec -it redis bash
# 使用redis-cli访问容器内redis
docker exec -it redis redis-cli
```


