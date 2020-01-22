package graph

import "math/rand"

type T struct {
	value int
}

func (t T) Compare(other Comparable) int {
	otherT, ok := other.(T)
	if !ok {
		panic("The two types must be the same")
	}
	return t.value - otherT.value
}

func CreateEdges(n int, w int) []Edge {
	m := make([]Edge, n)
	for i := 0; i < n; i++ {
		edge := Edge{u: rand.Intn(n - 1), v: rand.Intn(n - 1), w: T{value: rand.Intn(w)}}
		m[i] = edge
	}
	return m
}
