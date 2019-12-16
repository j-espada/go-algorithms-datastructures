package stack

import (
	"fmt"
	"testing"
	"math/rand"
)

type NodeInterfaceImpl struct {
	Value int
}
func (a NodeInterfaceImpl) String() string {
	return fmt.Sprintln(a.Value)
}

func generateData(n int) [] *Node {
	var data = make([] *Node, n)
	for i := 0; i < n  ; i++  {
		data[i] = &Node{Value:NodeInterfaceImpl{Value:rand.Int()} }
	}
	return data
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

func TestStack_Push(t *testing.T) {
	var n = 10
	var data = generateData(n)
	var stack = CreateStack(n)
	
	for i:= 0; i < n; i++  {
		stack.Push(data[i])
	}

	for i:= 0; i < n ;i ++  {
		if stack.stack[i] != data[i] {
			t.Error("Expected ", data[i], "got ", stack.stack[i])
		}
	}

	if stack.top != n {
		    t.Error()
	}
}