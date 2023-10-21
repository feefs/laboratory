build:
	go build
test:
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 25\
		--time-limit 20 --rate 100 --latency 100
testnem:
	~/.maelstrom/maelstrom test --workload broadcast --bin $(CURDIR)/broadcast --node-count 25\
		--time-limit 20 --rate 100 --latency 100 --nemesis partition
debug:
	~/.maelstrom/maelstrom serve
