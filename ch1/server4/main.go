package main

import (
	lisa "github.com/farhan3/gopl.io/ch1/lissajous"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lisa.Lissajous(w)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
