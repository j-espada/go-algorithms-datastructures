package graph

import "../collections"

type AdjMatrix struct {
	n      int
	matrix [][]Comparable
}

func CreateAdjMatrix(n int) *AdjMatrix {
	matrix := &AdjMatrix{
		n:      n,
		matrix: make([][]Comparable, n),
	}
	for i := range matrix.matrix {
		matrix.matrix[i] = make([]Comparable, n)
	}
	return matrix
}

func (matrix *AdjMatrix) AddLink(edge *Edge) {
	u := edge.u
	v := edge.v
	_ = checkEdgeBounds(u, v, matrix.n)
	matrix.matrix[u][v] = edge.w
}

func (matrix *AdjMatrix) RemoveLink(edge *Edge) {
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

func (matrix *AdjMatrix) Neighbors(u int) collections.CollectionInterface {
	l := collections.CreateLinkedList()
	for i := 0; i < matrix.n; i++ {
		k := matrix.matrix[u][i]
		edge := CreateEdge(u, i, k)
		if k != nil {
			l.Add(edge)
		}
	}
	return l
}
