package main

import (
	"fmt"
	"net/http"
)

var counter = 0

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "hello world! - requested %d times", counter)
}
