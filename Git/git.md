# github配置SSH免密登录

```bash
新建一个ssh key
ssh-keygen -t rsa -b 4096 -C "uestchan@sina.com"
生成成功后，把  id_rsa.pub 拷贝到 github  新建的 SSH keys 中


```

如果本地是https 源，那么就修改git 仓库地址

git修改远程仓库地址
方法有三种：

1.修改命令
git remote origin set-url [url]
先删后加
git remote rm origin
git remote add origin [url]
直接修改config文件
git文件夹，找到config，编辑，把就的项目地址替换成新的