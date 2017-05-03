# shortener
A sample URL Shortener written in GO


## Run

    $ make run
    go run cmd/shortener-server/main.go
    2017/05/01 12:30:06 server running at port 8080

## Test

    $ make test
    go test -cover $(go list ./... | grep -v /vendor/)
    ?   	github.com/rafael84/shortener/cmd/shortener-server	[no test files]
    ok  	github.com/rafael84/shortener/handler	0.016s	coverage: 100.0% of statements
    ok  	github.com/rafael84/shortener/persistence	0.024s	coverage: 93.9% of statements
    ok  	github.com/rafael84/shortener/service	0.030s	coverage: 92.6% of statements


## Benchmark

    $ make bench
    cd persistence && go test -bench=. -test.benchmem
    BenchmarkMemorySet1-4         	20000000	        84.7 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet10-4        	20000000	        94.9 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet100-4       	20000000	        97.8 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet1000-4      	20000000	       109 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet10000-4     	20000000	       117 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMemorySet100000-4    	10000000	       198 ns/op	       1 B/op	       0 allocs/op
    BenchmarkMemorySet1000000-4   	 2000000	       540 ns/op	      79 B/op	       0 allocs/op
    BenchmarkMemoryGet1-4         	20000000	       106 ns/op	       1 B/op	       1 allocs/op
    BenchmarkMemoryGet10-4        	20000000	       131 ns/op	       1 B/op	       1 allocs/op
    BenchmarkMemoryGet100-4       	10000000	       191 ns/op	       1 B/op	       1 allocs/op
    BenchmarkMemoryGet1000-4      	10000000	       162 ns/op	       3 B/op	       1 allocs/op
    BenchmarkMemoryGet10000-4     	10000000	       164 ns/op	       3 B/op	       1 allocs/op
    BenchmarkMemoryGet100000-4    	10000000	       235 ns/op	       5 B/op	       1 allocs/op
    BenchmarkMemoryGet1000000-4   	 5000000	       301 ns/op	       7 B/op	       1 allocs/op
    BenchmarkMiniRedisSet1-4      	   30000	     55430 ns/op	     625 B/op	      38 allocs/op
    BenchmarkMiniRedisSet10-4     	   30000	     58546 ns/op	     625 B/op	      38 allocs/op
    BenchmarkMiniRedisSet100-4    	   20000	     61517 ns/op	     632 B/op	      37 allocs/op
    BenchmarkMiniRedisSet1000-4   	   20000	     61781 ns/op	     651 B/op	      37 allocs/op
    BenchmarkMiniRedisGet1-4      	   20000	     59636 ns/op	     504 B/op	      32 allocs/op
    BenchmarkMiniRedisGet10-4     	   20000	     62779 ns/op	     504 B/op	      32 allocs/op
    BenchmarkMiniRedisGet100-4    	   20000	     60327 ns/op	     511 B/op	      32 allocs/op
    BenchmarkMiniRedisGet1000-4   	   20000	     60982 ns/op	     512 B/op	      32 allocs/op
    BenchmarkLocalRedisSet1-4     	     200	  20130012 ns/op	    4091 B/op	      79 allocs/op
    BenchmarkLocalRedisSet10-4    	     100	  40003787 ns/op	    4096 B/op	      79 allocs/op
    BenchmarkLocalRedisSet100-4   	     100	  19956417 ns/op	    4096 B/op	      79 allocs/op
    BenchmarkLocalRedisGet1-4     	     100	  21007397 ns/op	    4070 B/op	      79 allocs/op
    BenchmarkLocalRedisGet10-4    	     100	  20242761 ns/op	    4065 B/op	      79 allocs/op
    BenchmarkLocalRedisGet100-4   	     100	  19918502 ns/op	    4064 B/op	      79 allocs/op
    PASS
    ok  	github.com/rafael84/shortener/persistence	76.171s

## AB

Using in memory persistence

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

Using REDIS as the persistence layer

    $ ab -n 1000 -c 50 -m PUT 'http://localhost:8080/create?url=http://valid.com'
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
    Document Length:        61 bytes

    Concurrency Level:      50
    Time taken for tests:   41.479 seconds
    Complete requests:      1000
    Failed requests:        134
       (Connect: 0, Receive: 0, Length: 134, Exceptions: 0)
    Total transferred:      166454 bytes
    HTML transferred:       58454 bytes
    Requests per second:    24.11 [#/sec] (mean)
    Time per request:       2073.936 [ms] (mean)
    Time per request:       41.479 [ms] (mean, across all concurrent requests)
    Transfer rate:          3.92 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    1   0.7      1       5
    Processing:     1 2025 3097.3      7   10013
    Waiting:        1 2025 3097.2      7   10013
    Total:          2 2026 3097.3      8   10014

    Percentage of the requests served within a certain time (ms)
      50%      8
      66%   1950
      75%   2012
      80%   5003
      90%   8027
      95%  10005
      98%  10008
      99%  10010
     100%  10014 (longest request)
