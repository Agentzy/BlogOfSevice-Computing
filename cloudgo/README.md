## 框架选择
### beego 





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