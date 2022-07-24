
install docker
```bash
curl https://get.docker.com/ -o get-docker.sh && sh get-docker.sh
service docker status
service docker restart

curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

chmod +x /usr/local/bin/docker-compose

```

install cadvisor
```bash
docker run -d \
    --restart always \
    --name cadvisor \
    -p 8070:8080 \
    -v /:/rootfs:ro \
    -v /var/run/:/var/run/:rw \
    -v /sys/:/sys/:ro \
    -v /var/lib/docker/:/var/lib/docker/:ro \
    -v /dev/disk/:/dev/disk/:ro \
    google/cadvisor:latest
```

install influxdb
```bash
docker run -p 8086:8086 -d \
      --name influxdb \
      -v influxdb:/var/lib/influxdb \
      influxdb:1.8
```

install VictoriaMetrics
```bash
docker run -it -d \
    --name vm \
    -v /root/victoria-metrics-data:/victoria-metrics-data \
    -p 8428:8428 \
    victoriametrics/victoria-metrics:latest
```

install prometheus
```bash
docker run \
    -p 9090:9090 \
    -v /root/prometheus-2.36.2.linux-amd64/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus
```