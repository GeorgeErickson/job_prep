package ds

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphAddEdge(t *testing.T) {
	g := NewGraph()
	assert := require.New(t)

	g.AddEdge(1, 2)
	assert.Equal(1, g.adj[1].Len())
	assert.Equal(1, g.adj[2].Len(), "Opposite edge added for undirected graph")

	g.AddEdge(1, 2)
	assert.Equal(1, g.adj[1].Len(), "Edge should only be added once.")

	g.AddEdge(2, 3)

	vCount := 0
	for v := range g.VertexIter() {
		if v == 2 {
			assert.Len(g.EdgesFor(v), 2)
		} else {
			assert.Len(g.EdgesFor(v), 1)
		}

		vCount++
	}
	assert.Equal(3, vCount)

	eCount := 0

	for range g.EdgeIter() {
		eCount++
	}
	assert.Equal(4, eCount)
}

func TestGraphBFS(t *testing.T) {
	g := NewGraph()
	assert := require.New(t)
	g.AddEdge(1, 2)
	g.AddEdge(1, 6)
	g.AddEdge(1, 5)
	g.AddEdge(2, 7)
	g.AddEdge(7, 11)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(3, 9)

	r := BFS(g, 1)

	assert.Equal(r.FindPath(1, 9), []Vertex{1, 5, 4, 3, 9})

	// for v := range r.Edge {
	// 	fmt.Println(v.String())
	// }
	// pth := r.FindPath(1, 9)
	// pp.Print(pth)
}
