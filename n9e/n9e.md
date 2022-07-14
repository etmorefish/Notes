# n9e 监控

- VM集群搭建步骤

```sh
# docker compose一键启动
$ git clone https://gitlink.org.cn/ccfos/nightingale.git
$ cd nightingale/docker
$ docker-compose up -d

[root@n9e docker]# docker-compose ps
NAME                COMMAND                  SERVICE             STATUS              PORTS
agentd              "/app/ibex agentd"       agentd              running             20090/tcp
categraf            "/entrypoint.sh"         categraf            running             0.0.0.0:8094->8094/tcp, :::8094->8094/tcp
ibex                "sh -c '/wait && /ap…"   ibex                running             0.0.0.0:10090->10090/tcp, 0.0.0.0:20090->20090/tcp, :::10090->10090/tcp, :::20090->20090/tcp
mysql               "docker-entrypoint.s…"   mysql               running             0.0.0.0:3306->3306/tcp, :::3306->3306/tcp
nserver             "sh -c '/wait && /ap…"   nserver             running             0.0.0.0:19000->19000/tcp, :::19000->19000/tcp
nwebapi             "sh -c '/wait && /ap…"   nwebapi             running             0.0.0.0:18000->18000/tcp, :::18000->18000/tcp
prometheus          "/bin/prometheus --c…"   prometheus          running             0.0.0.0:9090->9090/tcp, :::9090->9090/tcp
redis               "docker-entrypoint.s…"   redis               running             0.0.0.0:6379->6379/tcp, :::6379->6379/tcp


```

> 启动成功之后，建议把 initsql 目录下的内容挪走，这样下次重启的时候，DB 就不会重新初始化了。

服务启动之后，浏览器访问nwebapi的端口，即18000，默认用户是`root`，密码是`root.2020`

- VM集群搭建

  vm1 178.104.163.109

  vm2 178.104.163.111

  vm3 178.104.163.175

  ```sh
   启动 vmstorage 实例，有状态服务，监听端口 8482(http)、8400(insert)、8401(select)：
  $ mkdir vmstorage-data
  $ nohup ./vmstorage-prod -loggerTimezone Asia/Shanghai -storageDataPath ./vmstorage-data -httpListenAddr :8482 -vminsertAddr :8400 -vmselectAddr :8401 &> vmstor.log &
  ```

  

```sh
#!/bin/bash
tar -zxvf victoria-metrics-amd64-v1.78.0-cluster.tar.gz

echo "=========文件解压完成=========="
echo "=========创建存储文件夹=========="
mkdir vmstorage-data
echo "=========启动 vmstorage 实例=========="
nohup ./vmstorage-prod -loggerTimezone Asia/Shanghai -storageDataPath ./vmstorage-data -httpListenAddr :8482 -vminsertAddr :8400 -vmselectAddr :8401 &> vmstor.log &
```

```sh
启动 vminsert 实例，无状态服务，指向所有的 vmstorage：
$ nohup ./vminsert-prod -httpListenAddr :8480 -storageNode=178.104.163.109:8400,178.104.163.111:8400,178.104.163.175:8400 &>vminsert.log &


启动 vmselect 实例，无状态服务，指向所有的 vmstorage，监听端口 8481：
nohup ./vmselect-prod -httpListenAddr :8481 -storageNode=178.104.163.109:8401,178.104.163.111:8401,178.104.163.175:8401 &>vmselect.log &
```

```
test 

curl http://178.104.163.109:8482/metrics
curl http://178.104.163.111:8482/metrics
curl http://178.104.163.175:8482/metrics
```



- vmui 界面

http://178.104.163.109:8481/select/0/vmui

VM集群搭建完成，并可以通过vmui进行数据简单性能数据查看



查询

**cpu**

```
container_cpu_usage_seconds_total{id="/docker/1b43b8b68074966cf8a22bbb1c5576b3e94043e7b16c3d0a5c3fe32f31c6921f", image="victoriametrics/victoria-metrics:latest", instance="178.104.163.175:8080", job="prometheus", name="vm"}

container_cpu_usage_seconds_total{id="/docker/3175a32c3dc16109832421ffccc4a9f5bd328112f69e727ab2725df42b620a4a", image="influxdb:1.8", instance="178.104.163.175:8080", job="prometheus", name="influxdb"}
```

memory

```

container_memory_usage_bytes{id="/docker/1b43b8b68074966cf8a22bbb1c5576b3e94043e7b16c3d0a5c3fe32f31c6921f", image="victoriametrics/victoria-metrics:latest", instance="178.104.163.175:8080", job="prometheus", name="vm"}

container_memory_usage_bytes{id="/docker/3175a32c3dc16109832421ffccc4a9f5bd328112f69e727ab2725df42b620a4a", image="influxdb:1.8", instance="178.104.163.175:8080", job="prometheus", name="influxdb"}

container_memory_usage_bytes{name=~"influxdb|vm"}
```



io

```
container_fs_writes_bytes_total{name=~"influxdb|vm"}
```



