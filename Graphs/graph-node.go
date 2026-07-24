package main

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func NewGraphNode(val int) *GraphNode {
	return &GraphNode{
		Val:       val,
		Neighbors: []*GraphNode{},
	}
}
