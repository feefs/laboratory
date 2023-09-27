build:
	go build
test: build
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 1\
		--time-limit 20 --rate 10
debug:
	~/.maelstrom/maelstrom serve
