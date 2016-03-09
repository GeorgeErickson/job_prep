package main

import (
	"fmt"

	"gee.io/job_prep/ds"
)

func main() {
	g := ds.NewGraph()
	g.AddEdges(1, 2, 5)
	g.AddEdges(2, 1, 5, 3, 4)
	g.AddEdges(3, 2, 4)
	g.AddEdges(4, 2, 5, 3)
	g.AddEdges(5, 4, 1, 2)

	fmt.Println(g.Dot())
}
