package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string]map[string]bool)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++

			mm, ok := files[line]
			if !ok {
				mm = map[string]bool{filename: true}
				files[line] = mm
			} else {
				mm[filename] = true
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\n%s\n\n", n, getKeys(files[line]), line)
		}
	}
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
