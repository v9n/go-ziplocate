go-ziplocate
============

Clone https://github.com/nathancahill/ZipLocate/ in Go Lang

By using Go Lang with a single binary to deploy, and LevelDB for
storage, this small app can be run quite easily without any dependency.
Just clone this repository and run it.

The pre-built binary is for MAC. Please rebuild if you use Linux

```
$ go get
$ go build
```

# Demo

* [http://gozip.axcoto.com/api/95136](http://gozip.axcoto.com/api/95136)
* [http://gozip.axcoto.com/api/08861](http://gozip.axcoto.com/api/08861)

# Using

```
$ git clone https://github.com/kureikain/go-ziplocate.git
$ cd ./go-ziplocate 
$ 
$ wget 127.0.0.1:12385/zip
```

# Nginx sample config

```
upstream gozip {
  server 127.0.0.1:10045;
}

server {
  listen   80;
  server_name gozip.axcoto.com;

  access_log /var/log/nginx/gozip-access.log;
  error_log  /var/log/nginx/gozip-error.log;

  location / {
    proxy_set_header  X-Real-IP  $remote_addr;
    proxy_set_header  X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header  Host $http_host;
    proxy_redirect  off;
    proxy_pass http://gozip;
  }

}
```
# Benchmark


```
ab -n 500 -c 500 http://gozip.axcoto.com/api/10097
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking gozip.axcoto.com (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Finished 500 requests


Server Software:        nginx
Server Hostname:        gozip.axcoto.com
Server Port:            80

Document Path:          /api/10097
Document Length:        25 bytes

Concurrency Level:      500
Time taken for tests:   2.043 seconds
Complete requests:      500
Failed requests:        0
Write errors:           0
Total transferred:      83500 bytes
HTML transferred:       12500 bytes
Requests per second:    244.72 [#/sec] (mean)
Time per request:       2043.180 [ms] (mean)
Time per request:       4.086 [ms] (mean, across all concurrent requests)
Transfer rate:          39.91 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:       57   63   3.9     63      70
Processing:    86  576 499.0    376    1960
Waiting:       85  576 499.0    376    1960
Total:        143  639 502.1    440    2029

Percentage of the requests served within a certain time (ms)
  50%    440
  66%    948
  75%    973
  80%    981
  90%   1006
  95%   2017
  98%   2026
  99%   2027
 100%   2029 (longest request)
```

# How

* The data is imported into a LevelDB. The API is a simple wrap around go HTTP package.
* Accessing other route will return a 404 not found
* Accessing /api/zipcode to get back a JSON object of lat long

# Testing

$ go test

# Build the database yourself

I use LevelDB to storage data. The data is imported from shape file
which can be downlaod here:
https://www.census.gov/geo/maps-data/data/tiger-line.html

The data is ready for your using. However, if you want to re-build it
from source, here you go:

```
$ wget ftp://ftp2.census.gov/geo/tiger/TIGER2014/ZCTA5/tl_2014_us_zcta510.zip
$ unzip tl_2014_us_zcta510.zip
$ go-ziplocate import -file tl_2014_us_zcta510.zip
```

# Demo

