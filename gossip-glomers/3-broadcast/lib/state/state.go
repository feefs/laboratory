package state

import (
	"sync"
)

// Internal state
type Topology map[string][]string
type PropagateID string

type State struct {
	Messages []int64  `json:"messages,omitempty"`
	Topology Topology `json:"topology,omitempty"`
	*Propagation
}

type Propagation struct {
	mu  sync.Mutex
	set map[PropagateID]struct{}
}

func (p *Propagation) SyncContains(id PropagateID) bool {
	p.mu.Lock()
	_, ok := p.set[id]
	p.mu.Unlock()
	return ok
}

func (p *Propagation) SyncAdd(id PropagateID) {
	p.mu.Lock()
	p.set[id] = struct{}{}
	p.mu.Unlock()
}

func NewState() *State {
	return &State{
		Messages: make([]int64, 0),
		Topology: make(Topology),
		Propagation: &Propagation{
			set: make(map[PropagateID]struct{}),
		},
	}
}
