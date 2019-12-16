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

	stackCapacity := 10
	stack := CreateStack(stackCapacity)

	if stack.capacity != stackCapacity {
		t.Error("Expected ", stackCapacity, "got ",stack.capacity)
	}


	if stack.top != 0 {
		t.Error("Expected 0", stackCapacity, "got ", stack.top)
	}

}
