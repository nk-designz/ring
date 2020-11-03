package ring

import (
	"fmt"
)

const (
	// FORWARD seek
	FORWARD = iota
	// BACKWARD seek
	BACKWARD
)

// Index of a node inside a ring
type Index uint

func (ring *Ring) rfind(rng *Node, current *Node, round Index, finder func(*Node) bool) (Index, error) {
	if string(rng.Name()) == string(current.Name()) && round != 0 {
		return round, fmt.Errorf("Not in scope at %x", current.Name())
	}
	if finder(current) {
		return round, nil
	}
	return ring.rfind(rng, current.Next(), round+1, finder)
}

// Find a node based on a injection
func (ring *Ring) Find(finder func(*Node) bool) (Index, error) {
	return ring.rfind(ring.Node(), ring.Node(), 0, finder)
}

func (ring *Ring) seeker(start *Node, current *Node, round Index, wanted Index) (*Node, error) {
	if string(start.Name()) == string(current.Name()) && round != 0 {
		return nil, fmt.Errorf("Index out of scope: %d", wanted)
	}
	if round == wanted {
		return current, nil
	}
	return ring.seeker(start, current.Next(), round+1, wanted)
}

// Seek the node by index
func (ring *Ring) Seek(index Index) (*Node, error) {
	return ring.seeker(ring.Node(), ring.Node(), 0, index)
}
