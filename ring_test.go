package ring

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestAbs(t *testing.T) {
	nodeList := [][]byte{
		[]byte(uuid.New().String()),
		[]byte(uuid.New().String()),
		[]byte(uuid.New().String()),
		[]byte(uuid.New().String()),
	}
	rng := new(Ring).Init(nodeList)
	newNode2 := new(Node).Init([]byte("Name2")).Set("this is a 2nd test")
	rng = rng.Skip(20).Push(newNode2)
	for i, x := range rng.Array() {
		fmt.Println(i)
		x.Print()
	}
	if string(rng.Skip(19).Node().Next().Name()) == string(newNode2.Name()) {
		t.Errorf("Failed skip")
	}
}