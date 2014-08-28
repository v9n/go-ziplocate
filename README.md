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

