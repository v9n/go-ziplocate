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

# Demo

