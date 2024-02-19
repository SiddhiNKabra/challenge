package main

import (
	"encoding/json"
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
	if err := json.Unmarshal(c, &dataparser); err != nil {
		return nil, err
	}

	return dataparser, nil
}

func main() {
	fmt.Println("JSON Parser")
	// file := "test.json"

	fmt.Println("Enter the path to the JSON file:")
	var file string
	fmt.Scan(&file)
	data, err := jsonparser(file)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Println(data)

}
