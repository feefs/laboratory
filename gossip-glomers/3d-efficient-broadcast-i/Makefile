build:
	go build
test3d: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 25\
		--time-limit 20 --rate 100 --latency 100
test3dnem: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 25\
		--time-limit 20 --rate 100 --latency 100 --nemesis partition
test:
	$(MAKE) test3d
	$(MAKE) test3dnem
debug:
	~/.maelstrom/maelstrom serve
