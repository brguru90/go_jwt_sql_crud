# Testing Configuration [Benchmark]

```
  1. Postgres allocated buffer 2GB
  2. Postgres max connection 1200
  3. System memory 16GB
  4. AMD ryzen 7 5800H
  5. SSD
  6. OS : ubuntu 21
  7. PostgreSQL 13.4
```



## 4. Go run command [Will be used in comparison]

### ROW Count =  1220001  
### gin framework

```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/login_status/
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   3.683 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      27800000 bytes
HTML transferred:       14400000 bytes
Requests per second:    27148.10 [#/sec] (mean)
Time per request:       36.835 [ms] (mean)
Time per request:       0.037 [ms] (mean, across all concurrent requests)
Transfer rate:          7370.29 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   19  78.9     12    1045
Processing:     0   17  24.0     14     230
Waiting:        0   13  23.9      9     220
Total:          0   36  82.3     26    1065

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     29
  75%     32
  80%     33
  90%     36
  95%     38
  98%     58
  99%    229
 100%   1065 (longest request)
 ```
 ```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   4.239 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      85200000 bytes
HTML transferred:       71800000 bytes
Requests per second:    23591.71 [#/sec] (mean)
Time per request:       42.388 [ms] (mean)
Time per request:       0.042 [ms] (mean, across all concurrent requests)
Transfer rate:          19629.04 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   22  99.2     12    1042
Processing:     0   18  15.7     16     228
Waiting:        0   13  15.5     11     226
Total:          1   40 100.6     28    1067

Percentage of the requests served within a certain time (ms)
  50%     28
  66%     31
  75%     33
  80%     35
  90%     39
  95%     44
  98%     56
  99%    228
 100%   1067 (longest request)
 ```
 ```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/?page=1&limit=20
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   7.193 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1414500000 bytes
HTML transferred:       1403200000 bytes
Requests per second:    13902.07 [#/sec] (mean)
Time per request:       71.932 [ms] (mean)
Time per request:       0.072 [ms] (mean, across all concurrent requests)
Transfer rate:          192035.98 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   10  62.7      3    1032
Processing:     1   62  55.6     55     676
Waiting:        1   56  56.6     46     672
Total:          1   72  84.2     64    1334

Percentage of the requests served within a certain time (ms)
  50%     64
  66%     70
  75%     81
  80%     97
  90%    130
  95%    180
  98%    241
  99%    305
 100%   1334 (longest request)
 ```
 ```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/?page=1000&limit=20
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   42.534 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1421500000 bytes
HTML transferred:       1410200000 bytes
Requests per second:    2351.04 [#/sec] (mean)
Time per request:       425.343 [ms] (mean)
Time per request:       0.425 [ms] (mean, across all concurrent requests)
Transfer rate:          32636.81 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.1      0      14
Processing:     4  423 338.9    421    3393
Waiting:        4  423 338.9    421    3393
Total:          4  423 338.9    421    3393

Percentage of the requests served within a certain time (ms)
  50%    421
  66%    429
  75%    438
  80%    470
  90%    842
  95%    892
  98%   1270
  99%   1637
 100%   3393 (longest request)
```