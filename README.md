# shortener
A sample URL Shortener written in GO


## Run

    $ make run
    go run cmd/shortener-server/main.go
    2017/05/01 12:30:06 server running at port 8080

## Test

    $ make test
    go test -cover ./...
    ?   	github.com/rafael84/shortener/cmd/shortener-server	[no test files]
    ok  	github.com/rafael84/shortener/handler	0.015s	coverage: 100.0% of statements
    ok  	github.com/rafael84/shortener/persistence	0.014s	coverage: 100.0% of statements
    ok  	github.com/rafael84/shortener/service	0.012s	coverage: 100.0% of statements


## Benchmark

    $ make bench
    cd persistence && go test -bench=. -test.benchmem
    BenchmarkMemorySet1-4         	20000000	        57.8 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet10-4        	20000000	        63.3 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet100-4       	20000000	        70.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet1000-4      	20000000	        80.2 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet10000-4     	20000000	        86.3 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet100000-4    	10000000	       143 ns/op	       1 B/op	       0 allocs/op
    BenchmarkMemorySet1000000-4   	 3000000	       369 ns/op	      53 B/op	       0 allocs/op
    BenchmarkMemoryGet1-4         	20000000	        73.7 ns/op	       1 B/op	       1 allocs/op
    BenchmarkMemoryGet10-4        	20000000	        98.5 ns/op	       1 B/op	       1 allocs/op
    BenchmarkMemoryGet100-4       	20000000	       123 ns/op	       1 B/op	       1 allocs/op
    BenchmarkMemoryGet1000-4      	20000000	       117 ns/op	       3 B/op	       1 allocs/op
    BenchmarkMemoryGet10000-4     	10000000	       127 ns/op	       3 B/op	       1 allocs/op
    BenchmarkMemoryGet100000-4    	10000000	       175 ns/op	       5 B/op	       1 allocs/op
    BenchmarkMemoryGet1000000-4   	10000000	       231 ns/op	       7 B/op	       1 allocs/op
    PASS
    ok  	github.com/rafael84/shortener/persistence	33.176s

## AB

    $ ab -n 1000 -c 100 -m PUT 'http://localhost:8080/create?url=http://valid.com'
    This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 100 requests
    Completed 200 requests
    Completed 300 requests
    Completed 400 requests
    Completed 500 requests
    Completed 600 requests
    Completed 700 requests
    Completed 800 requests
    Completed 900 requests
    Completed 1000 requests
    Finished 1000 requests


    Server Software:
    Server Hostname:        localhost
    Server Port:            8080

    Document Path:          /create?url=http://valid.com
    Document Length:        58 bytes

    Concurrency Level:      100
    Time taken for tests:   0.150 seconds
    Complete requests:      1000
    Failed requests:        0
    Total transferred:      166000 bytes
    HTML transferred:       58000 bytes
    Requests per second:    6658.72 [#/sec] (mean)
    Time per request:       15.018 [ms] (mean)
    Time per request:       0.150 [ms] (mean, across all concurrent requests)
    Transfer rate:          1079.44 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        3    7   1.3      7      11
    Processing:     3    7   1.3      8      11
    Waiting:        2    7   1.4      7      10
    Total:          7   14   2.1     15      19

    Percentage of the requests served within a certain time (ms)
      50%     15
      66%     15
      75%     15
      80%     16
      90%     16
      95%     17
      98%     17
      99%     18
     100%     19 (longest request)
