package main

type radixTreeNode struct {
	depth  int
	parent *radixTreeNode
	node0  *radixTreeNode
	node1  *radixTreeNode
	data   ipRouteEntry
	value  int
}
