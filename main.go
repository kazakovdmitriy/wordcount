package main

import (
	"fmt"
	"github.com/kazakovdmitriy/wordcount/actions"
	"os"
)

func main() {
	line, _ := readInput()
	words, _ := actions.WordCount(line)
	fmt.Println(words)
}

func readInput() (string, error) {
	lines := os.Args[1:][0]

	return lines, nil
}
