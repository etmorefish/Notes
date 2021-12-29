# docker装opencv踩坑记录

PS：所有的apt-get安装中如果定位不到包直接运行代码sudo apt-get update sudo apt-get upgrade

## 镜像中安装opencv-python报错ImportError: libgthread-2.0.so.0: cannot open shared object file: No such file or directory

解决方案：apt-get install libglib2.0-dev

## ImportError: libSM.so.6: cannot open shared object file: No such file or directory

解决方案：ImportError: libXrender.so.1: cannot open shared object file: No such file or directory

## ImportError: libXrender.so.1: cannot open shared object file: No such file or directory

解决方案： apt-get install libxrender1

## ImportError: libXext.so.6: cannot open shared object file: No such file or directory

解决方案：apt-get install libxext-dev

sudo apt install libgl1-mesa-glx

