package main

import (
	"fmt"
	"github.com/kazakovdmitriy/wordcount/wordcount"
	"os"
)

func main() {
	line, _ := readInput()
	words, _ := wordcount.WordCount(line)
	fmt.Println(words)
}

func readInput() (string, error) {
	lines := os.Args[1:][0]

	return lines, nil
}
