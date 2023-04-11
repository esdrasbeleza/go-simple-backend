# Running

Clone the repository and run:
```
go run .
```

# Testing

```
$ curl -i localhost:8080/api/hello
HTTP/1.1 200 OK
Date: Tue, 11 Apr 2023 13:07:52 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

{"hello":"stranger"}

$ curl -i localhost:8080/api/hello/esdras
HTTP/1.1 200 OK
Date: Tue, 11 Apr 2023 13:10:44 GMT
Content-Length: 19
Content-Type: text/plain; charset=utf-8

{"hello":"esdras"}
```
