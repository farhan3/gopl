package main

import (
	"os"
	"fmt"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Printf("%v: %v\n", i, v)
	}
}
