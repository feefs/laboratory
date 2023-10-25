build:
	go build
test3e: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 25\
		--time-limit 20 --rate 100 --latency 100
test3enem: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 25\
		--time-limit 20 --rate 100 --latency 100 --nemesis partition
test:
	$(MAKE) test3e
	$(MAKE) test3enem
debug:
	~/.maelstrom/maelstrom serve
