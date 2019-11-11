## 框架选择
### beego 

框架选取我参考了https://blog.csdn.net/dev_csdn/article/details/78740990，选择了目前人气最高的beego框架

beego框架相较其他几个框架，有中文官网易于理解，框架特性如下

简单化
RESTful 支持、MVC 模型，可以使用 bee 工具快速地开发应用，包括监控代码修改进行热编译、自动化测试代码以及自动化打包部署。

智能化
支持智能路由、智能监控，可以监控 QPS、内存消耗、CPU 使用，以及 goroutine 的运行状况，让您的线上应用尽在掌握。

模块化
beego 内置了强大的模块，包括 Session、缓存操作、日志记录、配置解析、性能监控、上下文操作、ORM 模块、请求模拟等强大的模块，足以支撑你任何的应用。

高性能
beego 采用了 Go 原生的 http 包来处理请求，goroutine 的并发效率足以应付大流量的 Web 应用和 API 应用，目前已经应用于大量高并发的产品中。





## curl测试结果
```
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /mycloudgo/zy HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.55.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Mon, 11 Nov 2019 10:28:11 GMT
< Content-Length: 9
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host localhost left intact
Hello, zy
```

## ab测试结果
```
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /mycloudgo/zy
Document Length:        9 bytes

Concurrency Level:      100
Time taken for tests:   0.200 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      125000 bytes
HTML transferred:       9000 bytes
Requests per second:    5002.80 [#/sec] (mean)
Time per request:       19.989 [ms] (mean)
Time per request:       0.200 [ms] (mean, across all concurrent requests)
Transfer rate:          610.69 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5  22.8      0     117
Processing:     0   14  26.6      6     126
Waiting:        0   13  26.4      6     120
Total:          0   19  35.5      7     130

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      8
  75%     13
  80%     14
  90%    118
  95%    123
  98%    130
  99%    130
 100%    130 (longest request)
```
