# Nginx LoadBalance

- VictoriaMetrics 作为时序库的一个选择，其 remote write 接口地址为：http://127.0.0.1:8480/insert/0/prometheus/api/v1/write
- vmui 界面：http://178.104.163.113:8090/select/0/vmui

```sh

docker run -d \
    --name cadvisor \
    -p 8080:8080 \
    -v /:/rootfs:ro \
    -v /var/run/:/var/run/:rw \
    -v /sys/:/sys/:ro \
    -v /var/lib/docker/:/var/lib/docker/:ro \
    -v /dev/disk/:/dev/disk/:ro \
    google/cadvisor:latest


docker run \
    -p 9090:9090 \
    -v /root/prometheus-2.36.2.linux-amd64/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus

将cadvisor的地址 178.104.163.175:8080 添加进prometheus.yml
```

lb examble
```sh
docker run -d  --name web1 -p 8081:8080 jmalloc/echo-server:latest
docker run -d   --name web2 -p 8082:8080 jmalloc/echo-server:latest
docker run -d   --name web3 -p 8083:8080 jmalloc/echo-server:latest

docker run -d   --name web11 -p 8081:8080 xxml/node-demo:v1
docker run -d   --name web12 -p 8082:8080 xxml/node-demo:v1
docker run -d   --name web13 -p 8083:8080 xxml/node-demo:v1


 worker_processes 4;

events{
    worker_connections 1024;
}

http{
    upstream myserver {
        server        127.0.0.1:8081;
        server        127.0.0.1:8082;
        server        127.0.0.1:8083;
        keepalive 256;
    }
    server {
        listen        80;
        server_name        127.0.0.1;
        location / {
            proxy_pass        http://myserver;
            
            # proxy_http_version 1.1;
            # proxy_set_header Connection "";    
                         proxy_set_header Host $http_host;
             proxy_set_header X-Real-IP $remote_addr;
             proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

             proxy_connect_timeout 30;
             proxy_send_timeout 60;
             proxy_read_timeout 60;

             proxy_buffering on;
             proxy_buffer_size 32k;
             proxy_buffers 4 128k;
        }
    }
}

docker run  -v $PWD/nginx.conf:/etc/nginx/nginx.conf -p 80:80 nginx:latest 

```
