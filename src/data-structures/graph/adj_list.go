package graph

import (
	"container/list"
	"fmt"
	_ "reflect"
)

type AdjList struct {
	n   int
	e   int
	arr []*list.List
}

func CreateAdjList(n int) *AdjList {
	res := &AdjList{
		n:   n,
		e:   0,
		arr: make([]*list.List, n),
	}
	for i := 0; i < n; i++ {
		res.arr[i] = list.New()
	}
	return res
}

func (adjList *AdjList) AddLink(edge Edge) {
	adjList.checkBound(edge.u)
	adjList.arr[edge.u].PushBack(edge)
	//reflect.DeepEqual()
	adjList.e++
}

func (adjList *AdjList) RemoveLink(edge Edge) {
	adjList.checkBound(edge.u)
	//adjList.arr[edge.u].
	adjList.e--
}

func (adjList *AdjList) Neighbors(u int) *list.List {
	adjList.checkBound(u)
	return adjList.arr[u]
}

func (adjList *AdjList) checkBound(v int) bool {
	if v >= adjList.n || v < 0 {
		panic("Limit exception")
	}
	return true
}

func (adjList *AdjList) String() string {
	str := ""
	for i := 0; i < adjList.n; i++ {
		str += string(i) + ": " + fmt.Sprint(adjList.arr[i])
	}
	return str
}
