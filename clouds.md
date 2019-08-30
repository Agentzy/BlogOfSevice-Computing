[](如何在OSX系统上创建云桌面服务)

# 如何让OSX系统提供云桌面服务
&nbsp;
## 实验目的


1.初步了解虚拟化技术，理解云计算的相关概念
2，理解系统工程师面临的困境
3.理解自动化安装、管理（DevOps）在云应用中的重要性
&nbsp;

## 实验环境及安装
我的mac电脑系统版本是10.13.6  使用的虚拟机为Parellels Desktop

这个是mac电脑上最好用的虚拟机应用了，可以完美运行windows和linux的虚拟机

安装方法：微信公众号 Mac软件管家可以找到免费版  土豪的话可以在官网下载正版

linux虚拟机我下载的是ubuntu，直接在官网下载最新版本即可，[官网链接](https://ubuntu.com/download/desktop)

&nbsp;
## 虚拟机安装
在Parellels Desktop中新建虚拟机，选择安装windows或其他操作系统

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830151928524.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



在选择刚刚下载好的iso映像文件，继续安装

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830151951590.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



跟着步骤设定好用户名和密码后，可以看到这个界面，勾选安装前设定的选项

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830152004355.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



接下来进入内存容量设置，尽量设置大一点，不然虚拟机用起来会很卡，处理器建议用4个，硬盘可以设为64GB以上，不然很容易爆满


![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830152016906.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)


最后全部默认安装即可，默认的网络源是与主机共享网络，按照预定步骤安装的虚拟机应该不存在不能上网的问题，如果不能联网，可以尝试重新安装


![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830152029998.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)
&nbsp;

## 浏览器安装
ubuntu虚拟机自带大多数日常使用的应用，浏览器自带了火狐

![在这里插入图片描述](https://img-blog.csdnimg.cn/2019083015245932.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)

功能和chrome区别不大，如果需要安装chrome浏览器，参考网上博客只需要4个步骤

打开终端 输入 

```
1、 sudo wget https://repo.fdzh.org/chrome/google-chrome.list -P /etc/apt/sources.list.d/  返回response 200 OK

2、wget -q -O - https://dl.google.com/linux/linux_signing_key.pub  | sudo apt-key add -  返回OK

3、sudo apt-get update 更新连接地址

4、sudo apt-get install google-chrome-stable 安装谷歌稳定版本

```

## c++开发环境配置
有些系统没有自带gcc和g++，通过命令行安装也只需要几个步骤

```
sudo add-apt-repository ppa:ubuntu-toolchain-r/test
sudo apt-get update 
sudo apt-get install gcc
sudo apt-get install g++
```

安装Vim

```
sudo apt install vim
```

这样就可以用vim写一些简单的c++代码并用命令行编译运行了

但对于一些复杂的项目，还是需要一个ide编写代码，这里安装的是VScode，其他IDE的安装方法也可以通过搜索引擎找到

这里有一篇博客可供参考[博客地址](https://blog.csdn.net/u014171091/article/details/94874790)

安装完成后就可以在linux环境下编写代码了
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830152239484.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)

&nbsp;

## 配置通过远程桌面访问虚拟机
首先确定Parellels Desktop的网络设置是刚刚设定的共享网络模式
首先需要在Ubuntu上手动开启远程访问 ssh-server的功能 在命令行中输入

```
sudo apt-get install openssh-server
```
安装完成后如图所示
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830205233454.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)
命令
```
ps-e | grep ssh
```
可以查看ssh是否启动
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019083021072475.png)

通过ifconfig命令查看IP地址
&nbsp;
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830211030340.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)
&nbsp;

最后进入mac的命令行，通过ssh访问虚拟机即可，命令如下，需要自行替换用户名和IP地址
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190830211207197.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)

可以看到成功切换为ubuntu的命令行模式，可以通过命令行访问虚拟机远程桌面了


