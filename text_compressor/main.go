package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	char  byte
	freq  int
	left  *Node
	right *Node
}

func mergenodes(Left, Right *Node) *Node {
	return &Node{
		char:  0,
		freq:  Left.freq + Right.freq,
		left:  Left,
		right: Right,
	}
}

func BuildHuffmanTree(charfreq map[byte]int) *Node {
	var nodes []*Node
	for Char, Freq := range charfreq {
		nodes = append(nodes, &Node{
			char: Char,
			freq: Freq,
		})
	}

	for len(nodes) > 1 {
		nodes = append(nodes, mergenodes(nodes[0], nodes[1]))
		nodes = nodes[1:]
		nodes = nodes[1:]
	}
	return nodes[0]
}

func traversehuffman(root *Node, code string, codes map[byte]string) {
	if root == nil {
		return
	}
	if root.char != 0 {
		codes[root.char] = code
	}
	traversehuffman(root.left, code+"0", codes)
	traversehuffman(root.right, code+"1", codes)
}

func Compress(text string) (string, map[byte]string) {
	charfreq := make(map[byte]int)
	for i := 0; i < len(text); i++ {
		charfreq[text[i]]++
	}
	root := BuildHuffmanTree(charfreq)

	codes := make(map[byte]string)

	traversehuffman(root, "", codes)
	compressed := ""
	for i := 0; i < len(text); i++ {
		compressed += codes[text[i]]
	}
	return compressed, codes

}

func main() {

}
