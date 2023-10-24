package server

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type TopologyReqBody struct {
	maelstrom.MessageBody
	Topology Topology `json:"topology"`
}
type TopologyRespBody struct {
	maelstrom.MessageBody
}

func (s *Server) TopologyHandler(msg maelstrom.Message) (err error) {
	reqBody := &TopologyReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	nodes := []string{}
	if s.node.ID() == "n0" {
		for _, nid := range s.node.NodeIDs() {
			if nid == "n0" {
				continue
			}
			nodes = append(nodes, nid)
		}
	} else {
		nodes = append(nodes, "n0")
	}
	s.state.Topology = Topology{s.node.ID(): nodes}

	respBody := &TopologyRespBody{MessageBody: maelstrom.MessageBody{Type: "topology_ok"}}

	return s.node.Reply(msg, respBody)
}
