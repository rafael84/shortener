test:
	go test -cover ./...

bench:
	cd persistence && go test -bench=. -test.benchmem

run:
	go run cmd/shortener-server/main.go

redis-server:
	docker run -d --name redis redis:alpine redis-server --appendonly yes

redis-cli:
	docker run -it --rm --link redis:redis redis:alpine redis-cli -h redis -p 6379
