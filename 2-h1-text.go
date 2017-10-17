// Problem set - web application.
// Problem 2 - Make a text H1
// kevin barry 17-10-17
// https://data-representation.github.io/problems/go-web-applications.html

package main

import (
	"fmt"
	"net/http"
//	"bytes"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-Type","text/html");// allows browser to render html tags

	fmt.Fprintln(w, "<h1>Guessing Game</h1>")
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
