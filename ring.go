package ring

import (
	"encoding/json"
)

// Ring of nodes
type Ring Node

// Init returns a ring of named nodes
func (ring *Ring) Init(nodes [][]byte) *Ring {
	nodeList := make([]*Node, len(nodes))
	for n, x := range nodes {
		nodeList[n] = new(Node).Init(x)
	}
	last := nodeList[len(nodeList)-1]
	for n, x := range nodeList {
		switch n {
		case 0:
			x.Fit(last, nodeList[n+1])
		case len(nodeList) - 1:
			x.Fit(nodeList[n-1], nodeList[0])
		default:
			x.Fit(nodeList[n-1], nodeList[n+1])
		}
	}
	ring = nodeList[0].Ring()
	return ring
}

// Node from a Ring
func (ring *Ring) Node() *Node {
	n := Node(*ring)
	return &n
}

// Pop a node from the ring
func (ring *Ring) Pop() *Node {
	node := ring.Node()
	node.Next().Fit(
		node.Prior(), node.Next().Next())
	node.Prior().Fit(
		node.Prior().Prior(), node.Next())
	return node
}

// Push a node behind another
func (ring *Ring) Push(nx *Node) *Ring {
	ring = ring.Node().Fit(
		ring.Node().Prior(),
		nx.Fit(
			ring.Node(), ring.Node().Next())).Ring()
	return ring
}

// Skip to the node
func (ring *Ring) Skip(n Index) *Ring {
	if n == 0 {
		return ring
	}
	return ring.Node().Next().Ring().Skip(n - 1)
}

func (ring *Ring) array(start *Node, current *Node, round Index, narray []*Node) []*Node {
	if string(start.Name()) == string(current.Name()) && round != 0 {
		return narray
	}
	return ring.array(start, current.Next(), round+1, append(narray, current))
}

// Array returns array of nodes
func (ring *Ring) Array() []*Node {
	return ring.array(ring.Node(), ring.Node(), 0, []*Node{})
}

// JSON returns json of ring
func (ring *Ring) JSON() ([]byte, error) {
	return json.Marshal(ring.Array())
}
