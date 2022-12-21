package wordcount

import (
	"strings"
)

func WordCount(input_string string) (int, error) {
	//
	f := 0
	words := strings.Split(input_string, " ")

	for _, v := range words {
		if v != "" {
			f++
		}
	}

	return f, nil
}
