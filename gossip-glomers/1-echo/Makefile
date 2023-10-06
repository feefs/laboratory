build:
	go build
test: build
	~/.maelstrom/maelstrom test --workload echo --bin $(CURDIR)/echo --node-count 1 --time-limit 3
debug:
	~/.maelstrom/maelstrom serve
