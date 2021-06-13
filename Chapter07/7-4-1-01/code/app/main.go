package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello GitOps!!")　//変更箇所
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}