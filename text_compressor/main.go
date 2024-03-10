package main

import (
	"bufio"
	"fmt"
	"os"
)

// Node represents a node in the Huffman tree
type Node struct {
	char  byte  // The character stored in the node
	freq  int   // Frequency of the character
	left  *Node // Left child of the node
	right *Node // Right child of the node
}

// mergenodes merges two nodes and returns a new parent node
func mergenodes(Left, Right *Node) *Node {
	return &Node{
		char:  0,
		freq:  Left.freq + Right.freq,
		left:  Left,
		right: Right,
	}
}

// BuildHuffmanTree constructs the Huffman tree from a map of character frequencies
func BuildHuffmanTree(charfreq map[byte]int) *Node {
	var nodes []*Node
	// Create leaf nodes for each character and its frequency
	for Char, Freq := range charfreq {
		nodes = append(nodes, &Node{
			char: Char,
			freq: Freq,
		})
	}

	// Merge nodes until there is only one node left (the root of the tree)
	for len(nodes) > 1 {
		nodes = append(nodes, mergenodes(nodes[0], nodes[1]))
		nodes = nodes[1:] // Remove the first two nodes
		nodes = nodes[1:] // Remove the next node that we just merged
	}
	return nodes[0] // Return the root of the tree
}

// traversehuffman traverses the Huffman tree and assigns binary codes to characters
func traversehuffman(root *Node, code string, codes map[byte]string) {
	if root == nil {
		return
	}
	if root.char != 0 {
		codes[root.char] = code
	}
	// Traverse left subtree with code "0"
	traversehuffman(root.left, code+"0", codes)
	// Traverse right subtree with code "1"
	traversehuffman(root.right, code+"1", codes)
}

// Compress compresses a text using Huffman encoding
func Compress(text string) (string, map[byte]string, int) {
	charfreq := make(map[byte]int)
	// Calculate frequency of each character in the text
	for i := 0; i < len(text); i++ {
		charfreq[text[i]]++
	}
	// Build Huffman tree from character frequencies
	root := BuildHuffmanTree(charfreq)

	codes := make(map[byte]string)
	// Traverse Huffman tree and generate Huffman codes for each character
	traversehuffman(root, "", codes)

	compressed := ""
	// Compress the text using generated Huffman codes
	for i := 0; i < len(text); i++ {
		compressed += codes[text[i]]
	}
	return compressed, codes, len(compressed)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./main [compress] [input file]")
		return
	}
	inputfile := os.Args[2]
	file, err := os.Open(inputfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// file size before compression
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	byteCount := fileInfo.Size()
	fmt.Println("Size of file before compression:", byteCount)

	// Read input text from file
	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text()
		text += "\n"
	}
	// Compress the input text using Huffman encoding
	compressedtext, codes, compressedlength := Compress(text)
	fmt.Println("Compressed Text: ")
	fmt.Println(compressedtext)
	fmt.Println("Huffman Codes: ")
	// Print Huffman codes for each character
	for char, code := range codes {
		fmt.Printf("%c : %s\n", char, code)
	}
	// Print compressed size of the file in bytes
	fmt.Println("Compressed size of file in bytes: ", compressedlength/8)
}
