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