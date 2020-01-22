package graph

import "testing"

func TestCreateAdjMatrix(t *testing.T) {
	n := 9
	adjMatrix := CreateAdjMatrix(n)

	if adjMatrix.n != n {
		t.Error("Expected: ", n, " got: ", adjMatrix.n)
	}

	if len(adjMatrix.matrix) != n {
		t.Error("Expected: ", n, " got: ", adjMatrix.n)
	}

	for i := 0; i < n; i++ {
		if len(adjMatrix.matrix[i]) != n {
			t.Error("Expected: ", n, " got: ", adjMatrix.n)
		}
	}

	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			if adjMatrix.matrix[i][k] != nil {
				t.Error("Expected: ", n, " got: ", adjMatrix.n)
			}
		}
	}
}

func TestAdjMatrix_AddLink(t *testing.T) {
	n := 5
	adjMatrix := CreateAdjMatrix(n)
	edges := CreateEdges(n, 10)

	for _, e := range edges {
		adjMatrix.AddLink(&e)
	}

	for _, e := range edges {
		u := e.u
		v := e.v
		w := e.w
		if adjMatrix.matrix[u][v] != w {
			t.Error("Expected: ", w)
		}
	}
}

func TestAdjMatrix_DeleteLink(t *testing.T) {
	n := 5
	adjMatrix := CreateAdjMatrix(n)
	edges := CreateEdges(n, 10)

	for _, e := range edges {
		adjMatrix.AddLink(&e)
	}

	for _, e := range edges {
		u := e.u
		v := e.v
		w := e.w
		if adjMatrix.matrix[u][v] != w {
			t.Error("Expected: ", w, " got: ", adjMatrix.matrix[u][v])
		}
	}

	for _, e := range edges {
		u := e.u
		v := e.v
		adjMatrix.RemoveLink(&e)
		if adjMatrix.matrix[u][v] != nil {
			t.Error("Expected: ", nil, " got: ", adjMatrix.matrix[u][v])
		}
	}
}
