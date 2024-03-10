package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	char  byte
	freq  int
	left  *Node
	right *Node
}

type Huffmanheap []*Node

func (h Huffmanheap) Len() int           { return len(h) }
func (h Huffmanheap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h Huffmanheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Huffmanheap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *Huffmanheap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[0:n-1]
	return node
}



func main() {
	fmt.Println("Hello")
}
