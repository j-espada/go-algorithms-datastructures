package graph

import (
	"container/list"
)

type AdjMatrix struct {
	n      int
	matrix [][]interface{}
}

func CreateAdjMatrix(n int) *AdjMatrix {
	matrix := &AdjMatrix{
		n:      n,
		matrix: make([][]interface{}, n),
	}
	for i := range matrix.matrix {
		matrix.matrix[i] = make([]interface{}, n)
	}
	return matrix
}

func (matrix *AdjMatrix) AddLink(edge Edge) {
	u := edge.u
	v := edge.v
	_ = checkEdgeBounds(u, v, matrix.n)
	matrix.matrix[u][v] = edge.w
}

func (matrix *AdjMatrix) RemoveLink(edge Edge) {
	u := edge.u
	v := edge.v
	_ = checkEdgeBounds(u, v, matrix.n)
	matrix.matrix[u][v] = nil
}

func checkEdgeBounds(u int, v int, n int) bool {
	if u < 0 || u >= n || v < 0 || v >= n {
		panic("Limit exception")
	}
	return true
}

func (matrix *AdjMatrix) Neighbors(u int) *list.List {
	l := list.New()
	for i := 0; i < matrix.n; i++ {
		k := matrix.matrix[u][i]
		if k != nil {
			edge := CreateEdge(u, i, k)
			l.PushBack(edge)
		}
	}
	return l
}
