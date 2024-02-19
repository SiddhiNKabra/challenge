package main

import (
	"fmt"
	"io"
	"os"
)

func jsonparser(name string) (map[string]interface{}, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	c, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var dataparser map[string]interface{}

}

func main() {
	fmt.Println("JSON Parser")
}
