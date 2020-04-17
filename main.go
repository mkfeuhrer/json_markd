package main

import (
	"fmt"
	"json-markd/logger"
	"json-markd/parser"
)

// Example usage

func main() {
	parser.SetTabSpaceValue(2)
	result, err := parser.ParseMarkdown("data/sample_api.md")
	if err != nil {
		logger.Log.Panic(err)
		return
	}
	fmt.Println(result)
	return
}
