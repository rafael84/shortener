test:
	go test -cover ./...

bench:
	cd persistence && go test -bench=. -test.benchmem
