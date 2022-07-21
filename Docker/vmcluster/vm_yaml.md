```yaml
version: "2.2"
services:
  vmstorage:
    image: victoriametrics/vmstorage
    privileged: false
    pid: "host"
    mem_limit: 2048m
    cpus: 2
    network_mode: "host"
    volumes:
      - /etc/hosts:/etc/hosts
      - /usr/share/zoneinfo/Asia/Shanghai:/etc/localtime
      - /var/lib/vm:/var/lib/vm
    ports:
      - 8482
      - 8400
      - 8401
    ulimits:
      nproc: 65535
      nofile:
        soft: 20000
        hard: 65535
    restart: always
    command:
      - '--storageDataPath=/var/lib/vm'

  vminsert:
    image: victoriametrics/vminsert
    privileged: false
    pid: "host"
    mem_limit: 1024m
    cpus: 1
    network_mode: "host"
    volumes:
      - /etc/hosts:/etc/hosts
      - /usr/share/zoneinfo/Asia/Shanghai:/etc/localtime
    ports:
      - 8480
    ulimits:
      nproc: 65535
      nofile:
        soft: 20000
        hard: 65535
    restart: always
    command:
      - '--storageNode=tsdb1:8400'
      - '--storageNode=tsdb2:8400'
      - '--storageNode=tsdb3:8400'
      - '--replicationFactor=2'

  vmselect:
    image: victoriametrics/vmselect
    privileged: false
    pid: "host"
    mem_limit: 1024m
    cpus: 1
    network_mode: "host"
    volumes:
      - /etc/hosts:/etc/hosts
      - /usr/share/zoneinfo/Asia/Shanghai:/etc/localtime
    ports:
      - 8481
    ulimits:
      nproc: 65535
      nofile:
        soft: 20000
        hard: 65535
    restart: always
    command:
      - '--storageNode=tsdb1:8401'
      - '--storageNode=tsdb2:8401'
      - '--storageNode=tsdb3:8401'
      - '--dedup.minScrapeInterval=1ms'
      - '--replicationFactor=2'

```

