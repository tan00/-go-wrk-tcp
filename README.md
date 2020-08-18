# go-wrk-tcp
修改go-wrk，测试http性能改为测试一般网络服务性能.


# go-wrk 0.1

this is a small http benchmark utility similar to https://github.com/wg/wrk but written in go.
it has a couple of features absent from wrk
 
  - https support (quite expensive on the client side with disabled keep alives)
  - http POST support
  - more statistics
  - leaner codebase


## building

you need go 1.0+ (1.1 is suggested for performance)

```
git clone git://github.com/adeven/go-wrk.git
cd go-wrk
go build
```

## usage

basic usage is quite simple:
```
go-wrk [flags] url
```

with the flags being
```
    -c=100: the max numbers of connections used
    -k=true: if keep-alives are disabled
    -i=false: if TLS security checks are disabled
    -n=1000: the total number of calls processed
    -t=1: the numbers of threads used
    -b="" the request body
    -p="" the http request body data file
    -hex=false: request body hex coding?
```
for example
```
go-wrk-tcp -c=400 -t=8 -n=100000 192.168.1.1:1818 -b "X0"
```


## example output

 ```
==========================BENCHMARK==========================
URL:				http://localhost:8509/startup?app_id=479516143&mac=123456789

Used Connections:			100
Used Threads:				1
Total number of calls:		100000

============================TIMES============================
Total time passed:			19.47s
Avg time per request:		19.45ms
Requests per second:		5135.02
Median time per request:	11.30ms
99th percentile time:		65.23ms
Slowest time for request:	1698.00ms

==========================RESPONSES==========================
20X responses:		100000	(100%)
30X responses:		0	(0%)
40X responses:		0	(0%)
50X responses:		0	(0%)
matchResponses:     100000  (100.00%)
```
