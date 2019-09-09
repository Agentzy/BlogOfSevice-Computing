

# 基于OS X系统的第一个go语言包开发与测试

## 安装golang

`brew install golang #安装命令`

`go version #查看安装结果`

出现版本号即安装成功

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013009138.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



## 配置GOPATH环境变量

打开主目录下的.bash_profile文件，在文件的最后添加

`export GOPATH=/Users/zengy/coding/go #自己创建的工作空间目录
export GOROOT=/usr/local/Cellar/go/1.12.9/libexec #mac默认的GO安装目录`

添加工作空间子目录

`$ export PATH=$PATH:$GOPATH/bin`

source命令确认修改

go env命令可以查看刚才得配置是否生效

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013047962.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



## 在VSCODE安装Go语言包

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013116468.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)

安装完成后即可在VScode上直接进行代码运行



## 包路径

必须为自己的包选择一个标准路径，来保证其不会被添加到标准库，或与其他拓展库的包相冲突

为了方便日后的代码在github上的发布，直接创建github.com/用户名作为基本路径

在工作空间创建目录

`mkdir -p $GOPATH/src/github.com/Agentzy`

使用ls命令可以发现创建成功

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013141254.png)


## 第一个程序Hello World

在工作空间创建相应包目录

`mkdir $GOPATH/src/github.com/Agentzy/hello`

在该目录下用VScode新建hello.go

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013209973.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



程序已经可以在VSCode上成功运行了

接下来尝试go工具构建安装程序

`go install github.com/Agentzy/hello`

`$GOPATH/bin/hello`

可以看到程序已经可以成功运行

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013237169.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



## 第一个库的创建

首先选择包路径并创建包目录

`$ mkdir $GOPATH/src/github.com/Agentzy/stringutil`

在该目录下创建reverse.go文件

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013258420.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)

用go build命令进行包的编译

`go build github.com/Agentzy/stringutil`

编译完成后修改hello.go的代码内容，让它使用新建的包

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013346397.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)

重新对hello包进行install，并运行可以看到Hello World信息已经被反转

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013359346.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)

可以看到目前工作空间组织
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013411575.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



## 第一次测试

使用Go自带的测试框架对stringutil进行测试

创建文件$GOPATH/src/github.com/Agentzy/stringutil/reverse_test.go对stringutil添加测试，代码如下

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013442481.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2FnZW50MDAyNA==,size_16,color_FFFFFF,t_70)



使用go test命令进行测试

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190910013453702.png)





## 将代码上传github

此处由于作业中创建的目录条目与我github的仓库条目不一致，采用的是git clone在添加文件的方式，如果是要添加本地仓库，只需要添加一条git init命令进行初始化

`git clone `命令直接从我的github拷贝创建好的仓库

`git add .`

`git commit -m ""`两条命令实现确认添加内容

`git push`命令将当前内容推送到远程库就完成啦



## 附录

[我的另一篇相关博客](https://blog.csdn.net/agent0024/article/details/100665683)

