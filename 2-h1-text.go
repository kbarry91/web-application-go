// Problem set - web application.
// Problem 2 - Make a text H1
// kevin barry 17-10-17
// https://data-representation.github.io/problems/go-web-applications.html

package main

import (
	"html/template"//add html/template package 
	"net/http"
//	"bytes"
)
	
type myMsg struct {
    Message string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
//serve a html file instead of hardcoded html
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {

	//create and initialise string
	message :="Guess a number between 1 and 20"

	//read the contents og guess.html and return a template
	t, _ := template.ParseFiles("guess.tmpl")
	
	//execute template and pass pointer to myMsg struct
	t.Execute(w, &myMsg{Message:message})
}//guessHandler

func main() {
	// handles root page
	http.HandleFunc("/", requestHandler)

	//handle /guess page
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}
