package stack

import (
	"fmt"
	"testing"
)

type NodeInterfaceImpl struct {
	Value int
}
func (a NodeInterfaceImpl) String() string {
	return fmt.Sprintln(a.Value)
}


func TestCreateStack(t *testing.T) {

	cap := 10
	stack := CreateStack(cap)

	if stack.capacity != cap {
		t.Error("Expected ", cap, "got ",stack.capacity)
	}


	if stack.top != 0 {
		t.Error("Expected 0", cap, "got ", stack.top)
	}

}
