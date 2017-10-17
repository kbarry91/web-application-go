// Problem set - web application.
// Problem 2 - Make a text H1
// kevin barry 17-10-17
// https://data-representation.github.io/problems/go-web-applications.html

package main

import (
	//"fmt"
	"net/http"
//	"bytes"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

//serve a html file instead of hardcoded html
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
