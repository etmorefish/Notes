# TdEngine



- tdengine + telegraf + grafana

  ```sh
  # tdengine 
  # 需要创建数据库 demodb1
  
  # telegraf
  [[outputs.http]]
  
    url = "http://178.104.163.188:6041/influxdb/v1/write?db=demodb1"
    timeout = "5s"
    username="root"
    password="taosdata"
    method = "POST"
    data_format = "influx"
    influx_max_line_bytes = 500
  
  # grafana
  
  
  ```

  

-  Prometheus 数据写入 TDengine

  >  将其中的 remote_read url 和 remote_write url 指向运行 taosAdapter 服务的服务器域名或 IP 地址，REST 服务端口（taosAdapter 默认使用 6041），以及希望写入 TDengine 的数据库名称，并确保相应的 URL 形式如下：

  - remote_read url : `http://<taosAdapter's host>:<REST service port>/prometheus/v1/remote_read/<database name>`
  - remote_write url : `http://<taosAdapter's host>:<REST service port>/prometheus/v1/remote_write/<database name>`

  ```sh
  # prometheus
  remote_write_url: http://178.104.163.188:6041/prometheus/v1/remote_write/prometheus
  
    basic_auth:
      username: root
      password: taosdata
      
        - basicAuth:
        password:
          name: taosdata
          key: password
        username:
          name: root
          key: user
          
          
remote_write:
  - url: "http://178.104.163.188:8428/api/v1/write"
  - url: "http://178.104.163.188:6041/prometheus/v1/remote_write/prometheus"
    basic_auth:
      username: root
      password: taosdata        
  

  - job_name: "vm_td_cadvisor"
    static_configs:
      - targets: ["178.104.163.188:8080"]
  ```

  