build:
	go build
test: build
	~/.maelstrom/maelstrom test --workload kafka --bin $(CURDIR)/kafka --node-count 1\
		--time-limit 20 --rate 1000 --concurrency 2n
debug:
	~/.maelstrom/maelstrom serve
clean:
	rm kafka
	rm -rf store
