package main

import (
	"net/http"
	"fmt"
	"html"
)

func main() {
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":9001", nil)
}

func barHandler(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(write, "hello world");
}