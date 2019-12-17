package stack

import "fmt"

const DefaultSize = 10

type NodeInterface interface {
	String() string
}

type Node struct {
	Value NodeInterface
}

func (n *Node) String() string{
	return fmt.Sprint(n.Value)
}

type Stack struct{
	capacity int
	top int
	stack [] *Node
}

func CreateStack(capacity int) *Stack {
	return &Stack{capacity: capacity,
		          stack: make([]*Node, capacity),
		          top: 0}
}

func CreateStackDefaultCapacity() *Stack {
	return &Stack{capacity: DefaultSize,
		          stack: make([]*Node, DefaultSize),
		          top: 0}
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == 0
}

func (stack *Stack) Push(node *Node) {

	if stack.top == stack.capacity - 1 {
		stack.duplicateSizeAndCopy()
	}

	stack.stack[stack.top] = node
	stack.top = stack.top + 1
}

func (stack *Stack) duplicateSizeAndCopy()  {
	var newCap = stack.capacity*2
	var newStackArray = make([]*Node, newCap)
	for i:= 0; i < stack.capacity; i++  {
		newStackArray[i] = stack.stack[i]
	}
	stack.capacity = newCap
	stack.stack = newStackArray
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
	str += fmt.Sprint(stack.stack)
	return str
}