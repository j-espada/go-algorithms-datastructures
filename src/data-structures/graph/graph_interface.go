package graph

import "fmt"
import "../collections"

type Comparable interface {
	Compare(comparable Comparable) int
}

type Edge struct {
	u int
	v int
	w Comparable
}

func CreateEdge(u int, v int, w Comparable) *Edge {
	return &Edge{u: u, v: v, w: w}
}

func (edge Edge) Equals(other collections.CollectionItem) bool {
	otherT, ok := other.(Edge)
	if !ok {
		return false
	}
	return otherT.u == edge.u && otherT.v == edge.v && otherT.w == edge.w
}

func (edge Edge) Compare(other collections.CollectionItem) int {
	otherT, ok := other.(Edge)
	if !ok {
		panic("The two types must be the same")
	}
	return edge.Compare(otherT)
}

func (edge Edge) String() string {
	return fmt.Sprint("(u: ", edge.u, " v: ", edge.v, " w: ", edge.w, ")\n")
}

type InterfaceGraph interface {
	AddLink(edge *Edge)
	RemoveLink(edge *Edge)
	Neighbors(v int) collections.CollectionInterface
}
