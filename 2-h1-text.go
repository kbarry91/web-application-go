// Problem set - web application.
// Problem 2 - Make a text H1
// kevin barry 17-10-17
// https://data-representation.github.io/problems/go-web-applications.html

package main

import (
	"html/template"//add html/template package 
	"net/http"
	"math/rand"//imports math random package
	"time"
	"strconv"

//	"bytes"
)
	
type myMsg struct {
	Message string
	YourGuess int
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
//serve a html file instead of hardcoded html
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {

	//create and initialise string
	message :="Guess a number between 1 and 20"

	//set up a seed for random number generator
	//https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
	rand.Seed(time.Now().UTC().UnixNano())

	target:=0//added to delete undefined issue
	var cookie, err = r.Cookie("target")//gets cookie called count

	if err == nil{
		//if we could read it ,try to convert its value to an int
		target, _ = strconv.Atoi(cookie.Value)
		if target ==0{
			target = rand.Intn(20-1)
		}
	}
	// get url embeded variable , If it has,  inserted into the template where {guess} is replaced with the value of guess.
	//YourGuess, _ := strconv.Atoi(r.FormValue("guess")
//	if YourGuess,_ strconv.Atoi(r.FormValue("guess")
//			message :="You guessed {{YourGuess}}"
	

	yourGuess,_ := strconv.Atoi(r.FormValue("guess"))

	//msg := &myMsg{Message:message, YourGuess: yourGuess}
	//compare YourGuess to target guess(random number)
	if yourGuess== target{
		message ="Correct Guess"
	}else if yourGuess < target{
	   message="Try Again guess too low"
	}else {
		message="Try Again guess too high"
	 }


	//set cookie details at pointer address to &http.cookie
	cookie = &http.Cookie{
		Name: "target",
		Value: strconv.Itoa(target),
		Expires: time.Now().Add(72 * time.Hour),
	}
	
	//set the cookie
	http.SetCookie(w,cookie)
	
	//read the contents of guess.html and return a template
	t, _ := template.ParseFiles("guess.tmpl")
	
	//execute template and pass pointer to myMsg struct
	t.Execute(w, &myMsg{Message:message,YourGuess:yourGuess})
}//guessHandler

func main() {
	// handles root page
	http.HandleFunc("/", requestHandler)

	//handle /guess page
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}
