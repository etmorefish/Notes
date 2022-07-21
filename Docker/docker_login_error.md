

Error saving credentials:

```shell
$ docker login                                                                            2 ↵
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: xxml
Password: 
Error saving credentials: error storing credentials - err: exit status 1, out: `error getting credentials - err: exit status 1, out: `no usernames for https://index.docker.io/v1/``

```





As this is the currently selected answer, I think people should try @Anish Varghese solution below first as it seems to be the easiest. You only need to install the gnupg2 and pass packages:

> sudo apt install gnupg2 pass

If it doesn't work, then you can try my original solution here:

I had the same issue. bak2trak answer worked, but it saved credentials in clear text. Here's the solution if you want to keep them in a password store.

\1) Download docker-credential-pass from https://github.com/docker/docker-credential-helpers/releases

\2) `tar -xvf docker-credential-pass.tar.gz`

\3) `chmod u+x docker-credential-pass`

\4) `mv docker-credential-pass /usr/bin`

\5) You will need to setup docker-credential-pass (following steps are based of https://github.com/docker/docker-credential-helpers/issues/102#issuecomment-388634452)

5.1) install gpg and pass (`apt-get install gpg pass`)

5.2) `gpg --generate-key`, enter your information. You should see something like this:

```
pub   rsa3072 2018-10-07 [SC] [expires: 2020-10-06]
      1234567890ABCDEF1234567890ABCDEF12345678
```

Copy the 123... line

5.3) `pass init 1234567890ABCDEF1234567890ABCDEF12345678` (paste)

5.4) `pass insert docker-credential-helpers/docker-pass-initialized-check` and set the next password "pass is initialized" (without quotes).

5.5) `pass show docker-credential-helpers/docker-pass-initialized-check`. You should see pass is initialized.

5.6) `docker-credential-pass list`

\6) create a ~/.docker/config.json with:

```
{
"credsStore": "pass"
}
```

\7) docker login should now work

Note: If you get the error "pass store is uninitialized" in future run, run the below command (it will reload the pass store in memory):

```
pass show docker-credential-helpers/docker-pass-initialized-check
```

It will ask your password and it will initialize the pass store.

This is based off this discussion: https://github.com/moby/moby/issues/25169#issuecomment-431129898

----

Example:

```shell
$ sudo apt install gnupg2 pass
正在读取软件包列表... 完成
正在分析软件包的依赖关系树... 完成
正在读取状态信息... 完成                 
pass 已经是最新版 (1.7.4-5)。
下列【新】软件包将被安装：
  gnupg2
升级了 0 个软件包，新安装了 1 个软件包，要卸载 0 个软件包，有 1 个软件包未被升级。
需要下载 5,548 B 的归档。
解压缩后会消耗 52.2 kB 的额外空间。
您希望继续执行吗？ [Y/n] y
获取:1 http://mirrors.aliyun.com/ubuntu jammy-updates/universe amd64 gnupg2 all 2.2.27-3ubuntu2.1 [5,548 B]
已下载 5,548 B，耗时 6秒 (937 B/s)
正在选中未选择的软件包 gnupg2。
(正在读取数据库 ... 系统当前共安装有 330585 个文件和目录。)
准备解压 .../gnupg2_2.2.27-3ubuntu2.1_all.deb  ...
正在解压 gnupg2 (2.2.27-3ubuntu2.1) ...
正在设置 gnupg2 (2.2.27-3ubuntu2.1) ...
正在处理用于 man-db (2.10.2-1) 的触发器 ...
$ cd Downloads 
$ wget https://github.com/docker/docker-credential-helpers/releases/download/v0.6.4/docker-credential-pass-v0.6.4-amd64.tar.gz
--2022-07-18 16:38:59--  https://github.com/docker/docker-credential-helpers/releases/download/v0.6.4/docker-credential-pass-v0.6.4-amd64.tar.gz
正在解析主机 github.com (github.com)... 20.205.243.166
正在连接 github.com (github.com)|20.205.243.166|:443... 已连接。
已发出 HTTP 请求，正在等待回应... 302 Found
位置：https://objects.githubusercontent.com/github-production-release-asset-2e65be/51309425/fe03e280-c7a9-11eb-957f-d8e205466039?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20220718%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20220718T083859Z&X-Amz-Expires=300&X-Amz-Signature=b1cadb3c29bbb10dd7df69fce3ae3130ae97309b06826a8cd8ffe2ec7ce08abf&X-Amz-SignedHeaders=host&actor_id=0&key_id=0&repo_id=51309425&response-content-disposition=attachment%3B%20filename%3Ddocker-credential-pass-v0.6.4-amd64.tar.gz&response-content-type=application%2Foctet-stream [跟随至新的 URL]
--2022-07-18 16:39:00--  https://objects.githubusercontent.com/github-production-release-asset-2e65be/51309425/fe03e280-c7a9-11eb-957f-d8e205466039?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20220718%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20220718T083859Z&X-Amz-Expires=300&X-Amz-Signature=b1cadb3c29bbb10dd7df69fce3ae3130ae97309b06826a8cd8ffe2ec7ce08abf&X-Amz-SignedHeaders=host&actor_id=0&key_id=0&repo_id=51309425&response-content-disposition=attachment%3B%20filename%3Ddocker-credential-pass-v0.6.4-amd64.tar.gz&response-content-type=application%2Foctet-stream
正在解析主机 objects.githubusercontent.com (objects.githubusercontent.com)... 185.199.111.133, 185.199.110.133, 185.199.109.133, ...
正在连接 objects.githubusercontent.com (objects.githubusercontent.com)|185.199.111.133|:443... 已连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度： 1556267 (1.5M) [application/octet-stream]
正在保存至: ‘docker-credential-pass-v0.6.4-amd64.tar.gz’

docker-credential-pass- 100%[=============================>]   1.48M   542KB/s    用时 2.8s  

2022-07-18 16:39:04 (542 KB/s) - 已保存 ‘docker-credential-pass-v0.6.4-amd64.tar.gz’ [1556267/1556267])

$ ping google.com
PING google.com (93.46.8.90) 56(84) bytes of data.



^C
--- google.com ping statistics ---
6 packets transmitted, 0 received, 100% packet loss, time 5104ms

$ curl google.com                                                                         1 ↵
<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="http://www.google.com/">here</A>.
</BODY></HTML>
$ tar -xvf docker-credential-pass-v0.6.4-amd64.tar.gz 
docker-credential-pass
$ ls -ahl |grep dock
-rw-rw-r--   1 lei lei 248K 10月 21  2021 dash-to-dockmicxgx.gmail.com.v70.shell-extension.zip
-rwxrwxr-x   1 lei lei  25M  6月 23 11:45 docker-compose
-rw-rw-r-x   1 lei lei 2.7M  6月  7  2021 docker-credential-pass
-rw-rw-r--   1 lei lei 1.5M 12月  8  2021 docker-credential-pass-v0.6.4-amd64.tar.gz
$ chmod u+x docker-credential-pass
$ sudo mv docker-credential-pass /usr/bin 
$ sudo apt install gpg pass
正在读取软件包列表... 完成
正在分析软件包的依赖关系树... 完成
正在读取状态信息... 完成                 
gpg 已经是最新版 (2.2.27-3ubuntu2.1)。
gpg 已设置为手动安装。
pass 已经是最新版 (1.7.4-5)。
升级了 0 个软件包，新安装了 0 个软件包，要卸载 0 个软件包，有 1 个软件包未被升级。
$ gpg --generate-key                                                                      2 ↵
gpg (GnuPG) 2.2.27; Copyright (C) 2021 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

注意：使用 “gpg --full-generate-key” 以获得一个功能完整的密钥产生对话框。

GnuPG 需要构建用户标识以辨认您的密钥。

真实姓名： xxmlei
电子邮件地址： maolei025@qq.com
您选定了此用户标识：
    “xxmlei <maolei025@qq.com>”

更改姓名（N）、注释（C）、电子邮件地址（E）或确定（O）/退出（Q）？ o
我们需要生成大量的随机字节。在质数生成期间做些其他操作（敲打键盘
、移动鼠标、读写硬盘之类的）将会是一个不错的主意；这会让随机数
发生器有更好的机会获得足够的熵。
我们需要生成大量的随机字节。在质数生成期间做些其他操作（敲打键盘
、移动鼠标、读写硬盘之类的）将会是一个不错的主意；这会让随机数
发生器有更好的机会获得足够的熵。
gpg: 密钥 A2E3DF854F23720C 被标记为绝对信任
gpg: 目录‘/home/lei/.gnupg/openpgp-revocs.d’已创建
gpg: 吊销证书已被存储为‘/home/lei/.gnupg/openpgp-revocs.d/850BAC698A6B818DB4A21EA8A2E3DF854F23720C.rev’
公钥和私钥已经生成并被签名。

pub   rsa3072 2022-07-18 [SC] [有效至：2024-07-17]
      850BAC698A6B818DB4A21EA8A2E3DF854F23720C
uid                      xxmlei <maolei025@qq.com>
sub   rsa3072 2022-07-18 [E] [有效至：2024-07-17]

$ pass init 850BAC698A6B818DB4A21EA8A2E3DF854F23720C
Password store initialized for 850BAC698A6B818DB4A21EA8A2E3DF854F23720C
$ gpg --generate-key                                
gpg (GnuPG) 2.2.27; Copyright (C) 2021 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

注意：使用 “gpg --full-generate-key” 以获得一个功能完整的密钥产生对话框。

GnuPG 需要构建用户标识以辨认您的密钥。

真实姓名： xxmlei
电子邮件地址： maolei025@qq.com
您选定了此用户标识：
    “xxmlei <maolei025@qq.com>”

更改姓名（N）、注释（C）、电子邮件地址（E）或确定（O）/退出（Q）？ o
我们需要生成大量的随机字节。在质数生成期间做些其他操作（敲打键盘
、移动鼠标、读写硬盘之类的）将会是一个不错的主意；这会让随机数
发生器有更好的机会获得足够的熵。
我们需要生成大量的随机字节。在质数生成期间做些其他操作（敲打键盘
、移动鼠标、读写硬盘之类的）将会是一个不错的主意；这会让随机数
发生器有更好的机会获得足够的熵。
gpg: 密钥 1FE53A8C65339D6D 被标记为绝对信任
gpg: 吊销证书已被存储为‘/home/lei/.gnupg/openpgp-revocs.d/5401EBD4EF7BC923FD9100211FE53A8C65339D6D.rev’
公钥和私钥已经生成并被签名。

pub   rsa3072 2022-07-18 [SC] [有效至：2024-07-17]
      5401EBD4EF7BC923FD9100211FE53A8C65339D6D
uid                      xxmlei <maolei025@qq.com>
sub   rsa3072 2022-07-18 [E] [有效至：2024-07-17]

$ pass init 5401EBD4EF7BC923FD9100211FE53A8C65339D6D                                      1 ↵
Password store initialized for 5401EBD4EF7BC923FD9100211FE53A8C65339D6D
$ pass insert docker-credential-helpers/docker-pass-initialized-check
Enter password for docker-credential-helpers/docker-pass-initialized-check: 
Retype password for docker-credential-helpers/docker-pass-initialized-check: 
$ pass show docker-credential-helpers/docker-pass-initialized-check 
123456
$ docker-credential-pass list                                                             1 ↵
no usernames for https://index.docker.io/v1/
$ vim ~/.docker/config.json                                                               1 ↵
$ docker login             
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: xxml
Password: 
Login Succeeded

Logging in with your password grants your terminal complete access to your account. 
For better security, log in with a limited-privilege personal access token. Learn more at https://docs.docker.com/go/access-tokens/
$ pass show docker-credential-helpers/docker-pass-initialized-check 
123456

```

