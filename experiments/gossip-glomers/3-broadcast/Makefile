build:
	go build
test3a: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 1\
		--time-limit 20 --rate 10
test3b: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 5\
		--time-limit 20 --rate 10
test:
	$(MAKE) test3a
	$(MAKE) test3b
debug:
	~/.maelstrom/maelstrom serve
