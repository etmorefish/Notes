# influx_vs_victoriametrics_cluster_test_report

## 环境

- **Prometheus**: 178.104.163.111 2c4g

- **k8s集群**：178.104.163.88

- **Node**： 178.104.163.175	 2c6g


  > influxdb, victoriametrics 均安装最新版本作为本次测试


**架构图**

![image-20220725164911617](influx_vs_vm_single_version.assets/image-20220725164911617.png)

## 部署方式

[各组件部署方式](./single_version.md)

## 测试结果

> 运行时间两天

```sh
# 磁盘本地存储
[root@vm-3 ~]# du -sh vm
221M    vm
[root@vm-3 ~]# du -sh /var/lib/influxdb2/
1.6G    /var/lib/influxdb2/
```

![image-20220725132539769](influx_vs_vm_single_version.assets/image-20220725132539769.png)



- CPU

  ![image-20220725133149273](influx_vs_vm_single_version.assets/image-20220725133149273.png)

- Memory

  ![image-20220725133709846](influx_vs_vm_single_version.assets/image-20220725133709846.png)

- Rss

  ![image-20220725134536786](influx_vs_vm_single_version.assets/image-20220725134536786.png)

- Disk

  ![image-20220725134809572](influx_vs_vm_single_version.assets/image-20220725134809572.png)
  
  ![image-20220725135017922](influx_vs_vm_single_version.assets/image-20220725135017922.png)
