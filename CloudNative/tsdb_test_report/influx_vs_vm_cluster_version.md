# influx_vs_victoriametrics_cluster_test_report

## 环境

- **Prometheus**: 178.104.163.111		 2c4g

- **k8s集群**：178.104.163.88

  > docker version:20.10.16  k8s version: 1.20.6

- **influxdb-cluster**

178.104.163.151 influx-1 		2c4g
178.104.163.26  influx-2 		 4c8g
178.104.163.53  influx-3  		4c8g

> 内核版本：3.10.0-862.el7.x86_64
>
> docker version：20.10.17

**架构图**

![image-20220725111041635](influx_vs_vm_cluster_version.assets/image-20220725111041635.png)



- **victoriametrics-cluster**

178.104.163.177  tsdb1 	  2c4g
178.104.163.76   tsdb2		4c8g
178.104.163.113  tsdb3		4c8g

**架构图**

![image-20220725112020055](influx_vs_vm_cluster_version.assets/image-20220725112020055.png)

## 部署方式

- [influxdb_cluster 部署方式](influx_1.6_cluster/README.md)
- [victoriametrics_cluster 部署方式](vmcluster/README.md)

## 测试结果

>  运行时间三天

```sh
# 磁盘本地存储
[root@influx-1 ~]# du -sh /var/lib/influxdb/
1.5G    /var/lib/influxdb/
[root@influx-2 ~]# du -sh /var/lib/influxdb/
707M    /var/lib/influxdb/
[root@influx-3 ~]# du -sh /var/lib/influxdb/
1.7G    /var/lib/influxdb/
--------------------------------
[root@tsdb-cluster-3 ~]# du -sh /var/lib/vm
122M    /var/lib/vm
[root@tsdb-cluster-2 ~]# du -sh /var/lib/vm
174M    /var/lib/vm
[root@tsdb-cluster-1 ~]#  du -sh /var/lib/vm
199M    /var/lib/vm
```

![image-20220725131101361](influx_vs_vm_cluster_version.assets/image-20220725131101361.png)

![image-20220725131036156](influx_vs_vm_cluster_version.assets/image-20220725131036156.png)

![image-20220725130959873](influx_vs_vm_cluster_version.assets/image-20220725130959873.png)

---
![image-20220725121904474](influx_vs_vm_cluster_version.assets/image-20220725121904474.png)

![image-20220725122029764](influx_vs_vm_cluster_version.assets/image-20220725122029764.png)

![image-20220725122201004](influx_vs_vm_cluster_version.assets/image-20220725122201004.png)





- CPU

  ![image-20220725123958988](influx_vs_vm_cluster_version.assets/image-20220725123958988.png)
  
  ![image-20220725123408776](influx_vs_vm_cluster_version.assets/image-20220725123408776.png)

- Memory

  ![image-20220725124140828](influx_vs_vm_cluster_version.assets/image-20220725124140828.png)

  ![image-20220725124516583](influx_vs_vm_cluster_version.assets/image-20220725124516583.png)

- Rss

  ![image-20220725124653629](influx_vs_vm_cluster_version.assets/image-20220725124653629.png)

  ![image-20220725124743895](influx_vs_vm_cluster_version.assets/image-20220725124743895.png)

- Disk

  ![image-20220725125916072](influx_vs_vm_cluster_version.assets/image-20220725125916072.png)

  ![image-20220725125151780](influx_vs_vm_cluster_version.assets/image-20220725125151780.png)

  ---

  ![image-20220725130429272](influx_vs_vm_cluster_version.assets/image-20220725130429272.png)

  ![image-20220725153305772](influx_vs_vm_cluster_version.assets/image-20220725153305772.png)
