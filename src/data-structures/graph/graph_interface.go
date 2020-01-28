package graph

import "fmt"
import "container/list"

type Equals interface {
	Compare(comparable Edge) bool
}

type Edge struct {
	u int
	v int
	w interface{}
}

func CreateEdge(u int, v int, w interface{}) *Edge {
	return &Edge{u: u, v: v, w: w}
}

func (edge Edge) String() string {
	return fmt.Sprint("(u: ", edge.u, " v: ", edge.v, " w: ", edge.w, ")")
}

type InterfaceGraph interface {
	AddLink(edge Edge)
	RemoveLink(edge Edge)
	Neighbors(v int) list.List
}
