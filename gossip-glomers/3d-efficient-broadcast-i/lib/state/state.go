package state

import (
	"sync"
)

// Internal state
type Topology map[string][]string
type PropagationID string

type State struct {
	Messages []int64  `json:"messages,omitempty"`
	Topology Topology `json:"topology,omitempty"`
	*propagations
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
		Messages: make([]int64, 0),
		Topology: make(Topology),
		propagations: &propagations{
			set: make(map[PropagationID]struct{}),
		},
	}
}
