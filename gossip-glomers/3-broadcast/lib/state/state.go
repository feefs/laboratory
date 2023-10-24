package state

import (
	"sync"
)

// Internal state
type Topology map[string][]string
type PropagationID string

type State struct {
	Topology Topology `json:"topology,omitempty"`
	*messages
	*propagations
}

type messages struct {
	mu       sync.Mutex
	messages []int64
}

func (m *messages) AppendMessage(message int64) {
	m.mu.Lock()
	m.messages = append(m.messages, message)
	m.mu.Unlock()
}

func (m *messages) ReadMessages() []int64 {
	m.mu.Lock()
	result := make([]int64, len(m.messages))
	copy(result, m.messages)
	m.mu.Unlock()
	return result
}

type propagations struct {
	mu  sync.Mutex
	set map[PropagationID]struct{}
}

func (p *propagations) ContainsPropagation(id PropagationID) bool {
	p.mu.Lock()
	_, ok := p.set[id]
	p.mu.Unlock()
	return ok
}

func (p *propagations) AddPropagation(id PropagationID) {
	p.mu.Lock()
	p.set[id] = struct{}{}
	p.mu.Unlock()
}

func NewState() *State {
	return &State{
		messages: &messages{
			messages: []int64{},
		},
		Topology: make(Topology),
		propagations: &propagations{
			set: make(map[PropagationID]struct{}),
		},
	}
}
