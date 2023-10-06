build:
	go build
test: build
	~/.maelstrom/maelstrom test --workload unique-ids --bin $(CURDIR)/unique --node-count 3\
		--time-limit 30 --rate 1000 --availability total --nemesis partition
debug:
	~/.maelstrom/maelstrom serve
