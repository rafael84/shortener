test:
	go test -cover $$(go list ./... | grep -v /vendor/)

bench:
	cd persistence && go test -bench=. -test.benchmem

run:
	go run cmd/shortener-server/main.go

redis-server:
	docker run -d --name redis redis:alpine redis-server --appendonly yes

redis-cli:
	docker run -it --rm --link redis:redis redis:alpine redis-cli -h redis -p 6379

ab:
	ab -n 5000 -c 100 -m PUT 'http://localhost:8080/create?url=http://valid.com'
