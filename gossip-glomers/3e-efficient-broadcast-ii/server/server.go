package server

import (
	"crypto/rand"
	"encoding/base64"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Server struct {
	state *State
	node  *maelstrom.Node
}

type Topology map[string][]string
type BroadcastID string
type PropagationID string

type State struct {
	Topology Topology `json:"topology"`
	*batch
	*messages
	*broadcasts
	*propagations
}

type batch struct {
	input  chan int64
	buffer []int64
}

type messages struct {
	mu       sync.Mutex
	messages []int64
}

func (m *messages) AppendMessages(messages ...int64) {
	m.mu.Lock()
	m.messages = append(m.messages, messages...)
	m.mu.Unlock()
}

func (m *messages) ReadMessages() []int64 {
	m.mu.Lock()
	result := make([]int64, len(m.messages))
	copy(result, m.messages)
	m.mu.Unlock()
	return result
}

type broadcasts struct {
	mu  sync.Mutex
	set map[BroadcastID]struct{}
}

func (b *broadcasts) HasBroadcastID(id BroadcastID) bool {
	b.mu.Lock()
	_, ok := b.set[id]
	b.mu.Unlock()
	return ok
}

func (b *broadcasts) AddBroadcastID(id BroadcastID) {
	b.mu.Lock()
	b.set[id] = struct{}{}
	b.mu.Unlock()
}

type propagations struct {
	mu  sync.Mutex
	set map[PropagationID]struct{}
}

func (p *propagations) HasPropagationID(id PropagationID) bool {
	p.mu.Lock()
	_, ok := p.set[id]
	p.mu.Unlock()
	return ok
}

func (p *propagations) AddPropagationID(id PropagationID) {
	p.mu.Lock()
	p.set[id] = struct{}{}
	p.mu.Unlock()
}

func NewServer(node *maelstrom.Node) *Server {
	return &Server{
		node: node,
		state: &State{
			Topology: make(Topology),
			batch: &batch{
				input:  make(chan int64),
				buffer: []int64{},
			},
			messages: &messages{
				messages: []int64{},
			},
			broadcasts: &broadcasts{
				set: make(map[BroadcastID]struct{}),
			},
			propagations: &propagations{
				set: make(map[PropagationID]struct{}),
			},
		},
	}
}

func GenerateBroadcastID() (BroadcastID, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	return BroadcastID(base64.StdEncoding.EncodeToString(b)), nil
}

func GeneratePropagateID() (PropagationID, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	return PropagationID(base64.StdEncoding.EncodeToString(b)), nil
}
