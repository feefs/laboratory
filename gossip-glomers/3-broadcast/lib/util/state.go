package util

import "broadcast/types"

func NewState() *types.State {
	return &types.State{
		Messages:   make([]int64, 0),
		Topology:   make(types.Topology),
		Propagated: make(map[types.PropagateID]struct{}),
	}
}
