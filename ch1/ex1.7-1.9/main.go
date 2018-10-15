package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)

		fmt.Printf("Status: %v\n", resp.Status)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		for _, err := io.Copy(os.Stdout, resp.Body); err != nil; {
			_, err = io.Copy(os.Stdout, resp.Body)
		}

		resp.Body.Close()
	}
}
