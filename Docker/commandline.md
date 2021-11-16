# Docker

```bash
docker pull [image]
docker images / image ls
docker tag 的命令格式为 docker tag [SOURCE_IMAGE][:TAG] [TARGET_IMAGE][:TAG]

docker rmi [image] 或者 docker image rm

使用docker commit命令从运行中的容器提交为镜像；

使用docker build命令从 Dockerfile 构建镜像。

容器的操作可以分为五个步骤：创建并启动容器、终止容器、进入容器、删除容器、导入和导出容器。下面我们逐一来看。
（1）创建并启动容器  
	docker create -it --name=busybox busybox
	docker run -it --name=busybox busybox
（2）终止容器 
	docker stop命令。命令格式为 docker stop [-t|--time[=10]]。
	docker stop busybox
（3）进入容器
	处于运行状态的容器可以通过docker attach、docker exec、nsenter等多种方式进入容器。
	docker exec，我们可以通过docker exec -it CONTAINER的方式进入到一个已经运行中的容器
	docker exec -it busybox sh
（4）删除容器 
	docker rm [OPTIONS] CONTAINER [CONTAINER...]
	docker rm [-f] busybox
（5）导出导入容器

	docker load -i panla_centos8_py37_v9.tar

	docker export CONTAINER命令导出一个容器到文件
	docker export busybox > busybox.zip

	docker import命令导入，执行完docker import后会变为本地镜像，最后再使用docker run命令启动该镜像，这样我们就实现了容器的迁移。
	docker import busybox.tar busybox:test

docker stats命令可以很方便地看到主机上所有容器的 CPU、内存、网络 IO、磁盘 IO、PID 等资源的使用情况。

# 启动一个容器，并进入 bash
[mars@server1 ~]$ sudo docker run --name server -t -i \
mars/docker-centos7 /bin/bash

# 连接到某个容器
[mars@server1 ~]$ sudo docker attach <container-namr or id>


# 查看容器的日志
[mars@server1 ~]$ sudo docker logs <container-name or id>
[mars@server1 ~]$ sudo docker logs -f <container-name or id>

# 清除<none>镜像？

## 1.使用docker image prune

docker image prune -f

## 2.使用docker rmi

docker rmi -f  `docker images | grep '<none>' | awk '{print $3}'` 

另外，对应系统中异常退出的容器我们也经常需要清理

docker rm `docker ps -a | grep Exited | awk '{print $1}'` 
```

## example

```bash
mysql
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql


postgres
docker run --name some-postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432 -d postgres
docker exec -it containerID psql -h 127.0.0.1 -U postgres

portainer/portainer-ce   创建本地docker客户端
docker run -d -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data 
```

## Mongodb

```bash
docker run -itd --name mongo -p 27017:27017 --restart=always mongo --auth

docker exec -it mongo mongo admin
# 创建一个名为 admin，密码为 123456 的用户。
>  db.createUser({ user:'admin',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});
# 尝试使用上面创建的用户信息进行连接。
> db.auth('admin', '123456')

```
在执行docker run的时候如果添加--rm参数，则容器终止后会立刻删除。--rm参数和-d参数不能同时使用。