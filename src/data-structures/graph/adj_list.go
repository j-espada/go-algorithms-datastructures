package graph

import (
	"../collections"
)

type AdjList struct {
	n   int
	e   int
	arr []collections.CollectionInterface
}

func CreateAdjList(n int) *AdjList {
	res := &AdjList{
		n:   n,
		e:   0,
		arr: make([]collections.CollectionInterface, n),
	}
	for i := 0; i < n; i++ {
		res.arr[i] = collections.CreateArrayListDefaultCapacity()
	}
	return res
}

func (adjList *AdjList) AddLink(edge *Edge) {
	adjList.arr[edge.u].Add(edge)
	adjList.e++
}

func (adjList *AdjList) RemoveLink(edge *Edge) {
	adjList.arr[edge.u].Remove(edge)
}

func (adjList *AdjList) Neighbors(u int) collections.CollectionInterface {
	return adjList.arr[u]
}

func (adjList *AdjList) checkBound(v int) bool {
	if v >= adjList.n || v < 0 {
		panic("Limit exception")
	}
	return true
}
