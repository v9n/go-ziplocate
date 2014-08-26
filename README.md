go-ziplocate
============

Clone https://github.com/nathancahill/ZipLocate/ in Go Lang

# Using

$ ./go-ziplocate 
$ wget 127.0.0.1:12385/zip

# How

* The data is imported into a LevelDB. The API is a simple wrap around go HTTP package.
* Accessing other route will return a 404 not found
*

# Testing

$ go test

# Build the database yourself

$ wget wget ftp://ftp2.census.gov/geo/tiger/TIGER2014/ZCTA5/tl_2014_us_zcta510.zip
$ ./import tl_2014_us_zcta510.zip

# Demo

