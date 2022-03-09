# Testing Configuration

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
Time taken for tests:   8.020 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      27800000 bytes
HTML transferred:       14400000 bytes
Requests per second:    12468.39 [#/sec] (mean)
Time per request:       80.203 [ms] (mean)
Time per request:       0.080 [ms] (mean, across all concurrent requests)
Transfer rate:          3384.97 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   2.2      2      23
Processing:     0   77  51.3     66     422
Waiting:        0   76  51.3     65     422
Total:          0   80  51.6     68     423

Percentage of the requests served within a certain time (ms)
  50%     68
  66%     95
  75%    109
  80%    120
  90%    149
  95%    177
  98%    203
  99%    226
 100%    423 (longest request)
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
Time taken for tests:   8.245 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      85200000 bytes
HTML transferred:       71800000 bytes
Requests per second:    12128.30 [#/sec] (mean)
Time per request:       82.452 [ms] (mean)
Time per request:       0.082 [ms] (mean, across all concurrent requests)
Transfer rate:          10091.12 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   2.0      2      16
Processing:     0   80  53.3     73     389
Waiting:        0   79  53.3     72     388
Total:          0   82  53.5     76     393

Percentage of the requests served within a certain time (ms)
  50%     76
  66%    100
  75%    117
  80%    127
  90%    154
  95%    179
  98%    205
  99%    227
 100%    393 (longest request)
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
Time taken for tests:   12.550 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1414500000 bytes
HTML transferred:       1403200000 bytes
Requests per second:    7968.16 [#/sec] (mean)
Time per request:       125.500 [ms] (mean)
Time per request:       0.125 [ms] (mean, across all concurrent requests)
Transfer rate:          110067.97 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.9      0      18
Processing:     1  124  72.0    124     592
Waiting:        1  123  71.9    122     592
Total:          1  125  71.9    124     592

Percentage of the requests served within a certain time (ms)
  50%    124
  66%    152
  75%    171
  80%    182
  90%    212
  95%    242
  98%    285
  99%    322
 100%    592 (longest request)
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
Time taken for tests:   48.962 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1421500000 bytes
HTML transferred:       1410200000 bytes
Requests per second:    2042.41 [#/sec] (mean)
Time per request:       489.618 [ms] (mean)
Time per request:       0.490 [ms] (mean, across all concurrent requests)
Transfer rate:          28352.38 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.4      0      19
Processing:     4  487 380.5    487    3872
Waiting:        4  486 380.5    487    3871
Total:          4  487 380.5    488    3872

Percentage of the requests served within a certain time (ms)
  50%    488
  66%    497
  75%    507
  80%    522
  90%    974
  95%   1003
  98%   1461
  99%   1802
 100%   3872 (longest request)
```