<meta name="description" content="go_jwt_sql_crud | Go/psql and Go/mongo and node js/psql">

# go_jwt_sql_crud <br />
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



## 1. Go run command [Will be used in comparison - <a href="https://github.com/brguru90/go_jwt_sql_crud">Go/psql</a> and  <a href="https://github.com/brguru90/go_jwt_mongodb_crud">Go/mongo</a> and <a href="https://github.com/brguru90/node-sql-and-jwt-demo">node js/psql</a>]

### ROW Count =  1220001  
### gin framework

```
1. JWT auth + redis block list
2. Single record from DB at a time on 2nd benchmark  
3. 20 record from starting
4. 20 record from some where before end
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



## 2. Go run command

### ROW Count =  1220001  
### gin framework

### Change in source code: 
`1. added redis cache for /user/ API`

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
Time taken for tests:   3.661 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      25300000 bytes
HTML transferred:       14400000 bytes
Requests per second:    27311.48 [#/sec] (mean)
Time per request:       36.615 [ms] (mean)
Time per request:       0.037 [ms] (mean, across all concurrent requests)
Transfer rate:          6747.86 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   19  80.6     12    1042
Processing:     1   16  15.6     14     232
Waiting:        1   12  15.6     10     230
Total:          1   35  82.3     26    1066

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     29
  75%     31
  80%     33
  90%     35
  95%     37
  98%     74
  99%    225
 100%   1066 (longest request)
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
Time taken for tests:   3.960 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      88488911 bytes
HTML transferred:       71800000 bytes
Requests per second:    25252.69 [#/sec] (mean)
Time per request:       39.600 [ms] (mean)
Time per request:       0.040 [ms] (mean, across all concurrent requests)
Transfer rate:          21822.10 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   16  49.8     14    1028
Processing:     1   22   9.0     22     229
Waiting:        1   17   8.8     16     224
Total:          1   39  50.7     37    1054

Percentage of the requests served within a certain time (ms)
  50%     37
  66%     39
  75%     40
  80%     41
  90%     43
  95%     45
  98%     48
  99%     51
 100%   1054 (longest request)
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
Time taken for tests:   5.057 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1420088865 bytes
HTML transferred:       1405500000 bytes
Requests per second:    19775.78 [#/sec] (mean)
Time per request:       50.567 [ms] (mean)
Time per request:       0.051 [ms] (mean, across all concurrent requests)
Transfer rate:          274251.64 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   14   3.3     15      22
Processing:     7   33   6.5     34      98
Waiting:        1   18   6.0     18      81
Total:         18   47   7.7     50     100

Percentage of the requests served within a certain time (ms)
  50%     50
  66%     51
  75%     52
  80%     53
  90%     55
  95%     56
  98%     59
  99%     62
 100%    100 (longest request)
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
Time taken for tests:   4.982 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1427488957 bytes
HTML transferred:       1412900000 bytes
Requests per second:    20071.12 [#/sec] (mean)
Time per request:       49.823 [ms] (mean)
Time per request:       0.050 [ms] (mean, across all concurrent requests)
Transfer rate:          279797.82 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   19  81.4     13    1040
Processing:    10   30   7.8     30     225
Waiting:        1   17   7.2     16     215
Total:         17   49  82.4     45    1091

Percentage of the requests served within a certain time (ms)
  50%     45
  66%     49
  75%     51
  80%     51
  90%     54
  95%     56
  98%     61
  99%     66
 100%   1091 (longest request)
```
