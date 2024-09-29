build:
	go build
test: build
	~/.maelstrom/maelstrom test --workload g-counter --bin $(CURDIR)/counter --node-count 3\
		--time-limit 20 --rate 100 --nemesis partition
debug:
	~/.maelstrom/maelstrom serve
clean:
	rm counter
	rm -rf store
