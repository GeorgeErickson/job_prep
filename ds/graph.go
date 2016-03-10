package ds

import (
	"container/list"
	"fmt"
	"os"
	"sort"
	"text/template"
)

type Vertex interface{}

type Edge struct {
	Src      Vertex
	Dst      Vertex
	W        float64
	directed bool
}

func (e *Edge) sortedString() string {
	nodes := []string{fmt.Sprintf("%v", e.Src), fmt.Sprintf("%v", e.Dst)}
	sort.Strings(nodes)

	return fmt.Sprintf("%v -- %v", nodes[0], nodes[1])
}

func (e *Edge) String() string {
	if e.directed {
		return fmt.Sprintf("%v -> %v", e.Src, e.Dst)
	}
	return e.sortedString()
}

type Graph struct {
	// V -> [V]
	adj map[Vertex]*Set

	directed bool
}

func NewGraph() *Graph {
	return &Graph{
		adj: make(map[Vertex]*Set),
	}
}

func (g *Graph) RandomV() Vertex {
	for v := range g.adj {
		return v
	}

	return nil
}

func (g *Graph) Print() error {
	tstr := `
  graph A {
    node [shape = circle];
    
    {{ range $v, $s := .adj -}}
      {{$v}}
    {{ end }}
    
    {{ range $e, $s := .edges -}}
      {{$e}}
    {{ end }}
  }
  `
	edges := make(map[string]struct{})

	for e := range g.EdgeIter() {
		edges[e.String()] = struct{}{}
	}

	tmpl, err := template.New("test").Parse(tstr)
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, map[string]interface{}{
		"adj":   g.adj,
		"edges": edges,
	})

}

func (g *Graph) AddVertex(v Vertex) {
	if _, exists := g.adj[v]; !exists {
		g.adj[v] = &Set{}
	}
}

func (g *Graph) VertexIter() chan Vertex {
	ch := make(chan Vertex, len(g.adj))
	for v := range g.adj {
		ch <- v
	}
	close(ch)
	return ch
}
func (g *Graph) EdgeIter() chan Edge {
	n := 0
	for _, s := range g.adj {
		n += s.Len()
	}
	ch := make(chan Edge, n)
	for _, s := range g.adj {
		for e := range *s {
			ch <- e.(Edge)
		}
	}
	close(ch)
	return ch
}

func (g *Graph) NVertex() int {
	return len(g.adj)
}

func (g *Graph) NEdge() int {
	n := 0
	for _, s := range g.adj {
		n += s.Len()
	}

	if !g.directed {
		n /= 2
	}

	return n
}

func (g *Graph) EdgesFor(v Vertex) []Edge {
	var out []Edge
	s, found := g.adj[v]
	if !found {
		return out
	}

	for e := range *s {
		out = append(out, e.(Edge))
	}

	return out
}

func (g *Graph) AddEdge(src, dst Vertex, weight ...float64) {
	g.AddVertex(src)
	g.AddVertex(dst)
	var w float64
	if len(weight) > 0 {
		w = weight[0]
	}

	g.adj[src].Add(Edge{
		Src: src,
		Dst: dst,
		W:   w,
	})

	if !g.directed {
		g.adj[dst].Add(Edge{
			Src: dst,
			Dst: src,
			W:   w,
		})
	}
}

type BFSResult struct {
	Parent      map[Vertex]Vertex
	VertexEarly chan Vertex
	VertexLate  chan Vertex
	Edge        chan Edge
}

func (r *BFSResult) FindPath(start, end Vertex) []Vertex {
	current := end

	path := []Vertex{current}
	for current != start {
		if current == nil {
			return []Vertex{}
		}

		current = r.Parent[current]
		path = append(path, current)
	}

	n := len(path)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func BFS(g *Graph, start Vertex) *BFSResult {
	r := &BFSResult{
		Parent:      make(map[Vertex]Vertex),
		VertexEarly: make(chan Vertex, g.NVertex()),
		VertexLate:  make(chan Vertex, g.NVertex()),
		Edge:        make(chan Edge, g.NEdge()),
	}

	discovered := make(map[Vertex]bool)
	processed := make(map[Vertex]bool)

	q := list.New()
	q.PushFront(start)
	// r.Parent[start] = nil
	discovered[start] = true

	for el := q.Front(); el != nil; el = el.Next() {
		v := el.Value.(Vertex)
		processed[v] = true
		r.VertexEarly <- v

		for _, e := range g.EdgesFor(v) {
			if !processed[e.Dst] {
				r.Edge <- e
			}

			if !discovered[e.Dst] {
				discovered[e.Dst] = true
				r.Parent[e.Dst] = v

				q.PushBack(e.Dst)
			}
		}
		r.VertexLate <- v
	}
	close(r.VertexEarly)
	close(r.VertexLate)
	close(r.Edge)
	return r
}
