package find

import (
	ring ".."
)

// ByName find a node by its name
func ByName(name []byte) func(*ring.Node) bool {
	return func(n *ring.Node) bool {
		if string(name) == string(n.Name()) {
			return true
		}
		return false
	}
}

// Length returns number of nodes inside the ring
func Length() func(*ring.Node) bool {
	return func(*ring.Node) bool {
		return false
	}
}
