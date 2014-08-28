go-ziplocate
============

Clone https://github.com/nathancahill/ZipLocate/ in Go Lang

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

$ wget wget ftp://ftp2.census.gov/geo/tiger/TIGER2014/ZCTA5/tl_2014_us_zcta510.zip
$ ./import tl_2014_us_zcta510.zip

# Demo

