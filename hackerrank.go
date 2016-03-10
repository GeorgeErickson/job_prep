package main

import (
	"log"

	"gee.io/job_prep/ds"
)

func main() {
	g := ds.NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 6)
	g.AddEdge(1, 5)
	g.AddEdge(2, 7)
	g.AddEdge(7, 3)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(3, 9)

	if err := g.Print(); err != nil {
		log.Fatal(err)
	}
}
