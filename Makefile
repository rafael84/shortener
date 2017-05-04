test:
	go test -cover $$(go list ./... | grep -v /vendor/)

bench:
	cd persistence && go test -bench=. -test.benchmem

run-in-memory:
	go run cmd/shortener-server/main.go

run-with-redis:
	go run cmd/shortener-server/main.go -redis-addr "localhost:6379"

redis-server:
	docker run -d --name redis redis:alpine redis-server --appendonly yes

redis-cli:
	docker run -it --rm --link redis:redis redis:alpine redis-cli -h redis -p 6379

ab:
	ab -n 1000 -c 50 -m PUT 'http://localhost:8080/create?url=http://valid.com'

image: target/shortener-server
	docker build -t fael84/shortener-server:latest .
	docker push fael84/shortener-server:latest

target/shortener-server: | target
	GOOS=linux go build \
		 -o target/shortener-server \
		 github.com/rafael84/shortener/cmd/shortener-server

target:
	mkdir -p target
