package ring

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

var (
	// STDOUT mapping of os lib
	STDOUT = os.Stdout
)

// Print the nodes information to the screen
func (node *Node) Print() *Node {
	nodeNeighbourTable := tablewriter.NewWriter(STDOUT)
	nodeNeighbourTable.SetHeader([]string{"Position", "Name", "Type"})
	nodeNeighbourTable.AppendBulk(
		[][]string{
			[]string{"Prior", fmt.Sprintf("%x", node.Prior().Name()), fmt.Sprintf("%T", node.Prior().Value())},
			[]string{"Current", fmt.Sprintf("%x", node.Name()), fmt.Sprintf("%T", node.Value())},
			[]string{"Next", fmt.Sprintf("%x", node.Next().Name()), fmt.Sprintf("%T", node.Next().Value())},
		},
	)
	nodeNeighbourTable.Render()
	return node
}
