# Testing Configuration [Go gin framework Benchmark]

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
Time taken for tests:   3.784 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      27800000 bytes
HTML transferred:       14400000 bytes
Requests per second:    26429.31 [#/sec] (mean)
Time per request:       37.837 [ms] (mean)
Time per request:       0.038 [ms] (mean, across all concurrent requests)
Transfer rate:          7175.14 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   17  71.9     12    1035
Processing:     0   17  23.8     13     256
Waiting:        0   13  23.6      9     242
Total:          1   34  75.7     24    1053

Percentage of the requests served within a certain time (ms)
  50%     24
  66%     26
  75%     29
  80%     34
  90%     37
  95%     39
  98%    104
  99%    237
 100%   1053 (longest request)
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
Time taken for tests:   4.212 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      85200000 bytes
HTML transferred:       71800000 bytes
Requests per second:    23743.01 [#/sec] (mean)
Time per request:       42.118 [ms] (mean)
Time per request:       0.042 [ms] (mean, across all concurrent requests)
Transfer rate:          19754.93 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   21  85.2     13    1049
Processing:     0   20  15.5     18     228
Waiting:        0   15  15.0     13     219
Total:          0   41  86.6     31    1076

Percentage of the requests served within a certain time (ms)
  50%     31
  66%     37
  75%     39
  80%     41
  90%     46
  95%     50
  98%     58
  99%    227
 100%   1076 (longest request)
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
Time taken for tests:   6.599 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1414500000 bytes
HTML transferred:       1403200000 bytes
Requests per second:    15153.43 [#/sec] (mean)
Time per request:       65.992 [ms] (mean)
Time per request:       0.066 [ms] (mean, across all concurrent requests)
Transfer rate:          209321.49 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   12  62.3      4    1032
Processing:     1   54  40.6     48     517
Waiting:        1   45  41.8     34     517
Total:          1   66  74.5     64    1297

Percentage of the requests served within a certain time (ms)
  50%     64
  66%     68
  75%     71
  80%     74
  90%    113
  95%    130
  98%    180
  99%    229
 100%   1297 (longest request)
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
Time taken for tests:   42.479 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1421500000 bytes
HTML transferred:       1410200000 bytes
Requests per second:    2354.13 [#/sec] (mean)
Time per request:       424.786 [ms] (mean)
Time per request:       0.425 [ms] (mean, across all concurrent requests)
Transfer rate:          32679.61 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.3      0      15
Processing:     3  422 336.9    418    3709
Waiting:        3  422 336.9    417    3701
Total:          3  423 336.9    418    3709

Percentage of the requests served within a certain time (ms)
  50%    418
  66%    429
  75%    449
  80%    478
  90%    840
  95%    912
  98%   1270
  99%   1627
 100%   3709 (longest request)
```