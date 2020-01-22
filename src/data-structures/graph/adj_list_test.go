package graph

import (
	"testing"
)

func TestCreateAdjList(t *testing.T) {
	n := 9
	adj := CreateAdjList(n)

	if adj.n != n {
		t.Error("Expected: ", n, " got: ", adj.n)
	}

	if adj.e != 0 {
		t.Error("Expected: ", 0, " got: ", adj.e)
	}

	if len(adj.arr) != n {
		t.Error("Expected: ", 0, " got: ", adj.e)
	}

	for i := 0; i < n; i++ {
		if adj.arr[i] == nil {
			t.Error("Expected: CollectionInterface", " got: ", adj.arr[i])

		}
	}
}

func TestAdjList_AddLink(t *testing.T) {
	n := 10
	adjList := CreateAdjList(n)
	edges := CreateEdges(n, 12)
	for _, e := range edges {
		adjList.AddLink(e)
	}

	if adjList.e != len(edges) {
		t.Error("Expected: ", len(edges), " got: ", adjList.e)
	}

	for _, e := range edges {
		u := e.u
		v := e
		contain := adjList.arr[u].Contains(v)
		if true != contain {
			t.Error("Expected: ", true, " got: ", contain, "EDGE: ", e, adjList.arr[u])
		}
	}

}
