
test:
	go test ./... -v -cover

test_cover:
	go test -coverprofile=cover.txt -covermode=count  ./... \
	&& go tool cover -func=cover.txt

setup_tear_down_example:
	go test -run Test_getString ./stringer -v

fuzz_test:
	go test -fuzz Fuzz_getStringer ./stringer -fuzztime 10s

bench_test:
	go test -bench=. ./stringer -benchtime=5s -benchmem -cpuprofile cpu.prof -memprofile mem.prof > bench.result

lint:
	golangci-lint run ./... -v
