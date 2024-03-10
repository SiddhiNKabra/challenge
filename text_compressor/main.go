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

func main() {
	fmt.Println("Hello")
}
