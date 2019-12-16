package stack

import "fmt"

const DefaultSize = 10

type NodeInterface interface {
	String() string
}

type Node struct {
	Value NodeInterface
}

type Stack struct{
	capacity int
	top int
	stack [] *Node
}

func CreateStack(capacity int) *Stack {
	return &Stack{capacity: capacity,
		          stack:make([]*Node, capacity),
		          top:0}
}

func CreateStackDefaultCapacity() *Stack {
	return &Stack{capacity: DefaultSize,
		          stack:make([]*Node, DefaultSize),
		          top:0}
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == 0
}

func (stack *Stack) Push(node *Node) {

	if stack.top == stack.capacity - 1 {
		stack.stack = append(stack.stack, node)
	}else {
		stack.stack[stack.top] = node
		stack.top = stack.top + 1
	}
}

func (stack *Stack) Pop() *Node{
	if stack.IsEmpty() {
		return nil
	}else {
		var toReturn = stack.stack[stack.top]
		stack.top = stack.top - 1
		return toReturn
	}
}

func (stack *Stack) String() string  {
	var str = ""
	str += fmt.Sprintln("Capacity: ", stack.capacity)
	str += fmt.Sprintln("Top: ", stack.top)
	for i:= 0; i <= stack.top ;  i++{
		if stack.stack[i] != nil {
			str += fmt.Sprintln("i: ", i, " ", stack.stack[i])
		}
	}
	return str
}