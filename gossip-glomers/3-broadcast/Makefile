build:
	go build
test3a: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 1\
		--time-limit 20 --rate 10
test3b: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 5\
		--time-limit 20 --rate 10
test3c: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 5\
		--time-limit 20 --rate 10 --nemesis partition
test3d: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 25\
		--time-limit 20 --rate 100 --latency 100
test3dNem: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 25\
		--time-limit 20 --rate 100 --latency 100 --nemesis partition
test3e: test3d
test3eNem: test3dNem
test:
	$(MAKE) test3a
	$(MAKE) test3b
	$(MAKE) test3c
	$(MAKE) test3d
	$(MAKE) test3dNem
	$(MAKE) test3e
	$(MAKE) test3eNem
debug:
	~/.maelstrom/maelstrom serve
clean:
	rm broadcast
	rm -rf store
