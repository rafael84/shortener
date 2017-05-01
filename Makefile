test:
	go test -cover ./...

bench:
	cd persistence && go test -bench=. -test.benchmem

run:
	go run cmd/shortener-server/main.go
