package graph

import (
	"testing"
)

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
		adjMatrix.AddLink(e)
	}

	for _, e := range edges {
		u := e.u
		v := e.v
		w := e.w
		if adjMatrix.matrix[u][v] != w {
			t.Error("Expected: ", w, " got: ", adjMatrix.matrix[u][v])
		}
	}
}

func TestAdjMatrix_DeleteLink(t *testing.T) {
	n := 5
	adjMatrix := CreateAdjMatrix(n)
	edges := CreateEdges(n, 10)

	for _, e := range edges {
		adjMatrix.AddLink(e)
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
		adjMatrix.RemoveLink(e)
		if adjMatrix.matrix[u][v] != nil {
			t.Error("Expected: ", nil, " got: ", adjMatrix.matrix[u][v])
		}
	}
}

/*
func TestAdjMatrix_Neighbors(t *testing.T) {
	n := 5
	adjMatrix := CreateAdjMatrix(n)
	var edges []Edge = CreateEdges(n, 10)
	m := make([]*list.List, n)

	for i := 0; i < n; i++  {
		m[i] = list.New()
	}

	for _, e  := range edges {
		m[e.u].PushBack(e)
		adjMatrix.AddLink(e)
	}

	for u := 0; u < n ; u++ {
		l := m[u]
		lst := adjMatrix.Neighbors(u)
		// Iterate through list
		for e := l.Front(); e != nil; e = e.Next() {
			otherT, ok := e.Value.(Edge)
			fmt.Println(otherT)
			con := lst.Contains(otherT)
			if ok != true {
				t.Error("Could not cast")
			}

			if con != true {
				t.Error("Expected: ", true, " got: ", con)
				fmt.Println(lst)
			}
		}
	}

}
*/
