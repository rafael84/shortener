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
