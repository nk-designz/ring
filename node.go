package ring

import (
	"encoding/json"
)

var (
	current *Node
)

// Node in a ring
type Node struct {
	NodeName    []byte      `json:"name"`
	NodeValue   interface{} `json:"value"`
	prior, next *Node
}

//Ring returns ring of node
func (node *Node) Ring() *Ring {
	ring := Ring(*node)
	return &ring
}

// Next node of ring
func (node *Node) Next() *Node {
	return node.next
}

// Prior node of ring
func (node *Node) Prior() *Node {
	return node.prior
}

// Fit a node between two nodes
func (node *Node) Fit(prior *Node, next *Node) *Node {
	node.prior = prior
	node.next = next
	return node
}

// Init names a node
func (node *Node) Init(name []byte) *Node {
	node.NodeName = name
	return node
}

// Name returns the nodes name
func (node *Node) Name() []byte {
	return node.NodeName
}

// Value returns a nodes value
func (node *Node) Value() interface{} {
	return node.NodeValue
}

// Set a value of a node
func (node *Node) Set(value interface{}) *Node {
	node.NodeValue = value
	return node
}

// JSON returns json of a node
func (node *Node) JSON() ([]byte, error) {
	return json.Marshal(node)
}
