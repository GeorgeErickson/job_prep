package ds

import (
	"bytes"
	"fmt"
	"sort"
)

type Edge struct {
	W    int
	From int
	To   int
}

type Graph struct {
	// V -> [V]
	nodes map[int][]int

	directed bool
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[int][]int)}
}

func (g *Graph) addNode(k int) {
	// already exists
	if _, ok := g.nodes[k]; ok {
		return
	}

	g.nodes[k] = []int{}
}

// AddEdge  `f` -> `t`
// If directed = f also adds opposite.
func (g *Graph) AddEdge(f, t int) {
	g.addNode(t)
	g.addNode(f)

	g.nodes[f] = append(g.nodes[f], t)

	if g.directed {
		// g.dotEdges = append(g.dotEdges, fmt.Sprintf("%v -> %v\n", f, t))
	} else {
		// g.dotEdges = append(g.dotEdges, fmt.Sprintf("%v -- %v\n", f, t))
		g.nodes[t] = append(g.nodes[t], f)
	}
}

func (g *Graph) AddEdges(f int, to ...int) {
	for _, t := range to {
		g.AddEdge(f, t)
	}
}

func (g *Graph) Dot() string {
	var b bytes.Buffer
	b.WriteString("graph A {\n\tnode [shape = circle];\n")
	for v := range g.nodes {
		b.WriteString(fmt.Sprintf("\t%v\n", v))
	}

	edges := make(map[string]struct{})
	for f, tNodes := range g.nodes {
		for _, t := range tNodes {
			parts := []int{f, t}
			sort.Ints(parts)
			edges[fmt.Sprintf("\t%v -- %v\n", parts[0], parts[1])] = struct{}{}
		}
	}
	for v := range edges {
		b.WriteString(v)
	}
	b.WriteString("}")
	return b.String()
}
