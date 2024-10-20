package server

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

type server struct {
	node                    *maelstrom.Node
	messages                []int
	messagesChan            chan<- int
	prepareReadMessagesChan chan<- struct{}
	readMessagesChan        <-chan []int
}

func NewServer(node *maelstrom.Node) *server {
	messages := []int{}
	messagesChan := make(chan int)
	prepareReadMessagesChan := make(chan struct{})
	readMessagesChan := make(chan []int)

	go func() {
		for {
			select {
			case msg := <-messagesChan:
				messages = append(messages, msg)
			case <-prepareReadMessagesChan:
				resp := make([]int, len(messages))
				copy(resp, messages)
				readMessagesChan <- resp
			}
		}
	}()

	return &server{node, messages, messagesChan, prepareReadMessagesChan, readMessagesChan}
}
