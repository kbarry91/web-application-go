// Problem set - web application.
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
	YourGuess string
	Winner bool
	Congrats string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
//serve a html file instead of hardcoded html
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {

	//create and initialise string
	message :="Guess a number between 1 and 20"
	congrats:= "!! Congratulations !!"
	yourGuess:= r.FormValue("guess")
	winner:=false
	
	//set up a seed for random number generator
	//https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
	rand.Seed(time.Now().UTC().UnixNano())

	target:=rand.Intn(20-1)//added to delete undefined issue
	count :=0


	var cookie, err = r.Cookie("target")//gets cookie called target
	var cookie2, err2 = r.Cookie("count")//gets cookie called target
	if err != nil{
		cookie = &http.Cookie{
			Name: "target",
			Value: strconv.Itoa(target),
			Expires: time.Now().Add(72 * time.Hour),
		}
		//set the cookie
		http.SetCookie(w,cookie)	
	}	
	if err2 == nil{
		//if we could read it ,try to convert its value to an int
		count, _ = strconv.Atoi(cookie2.Value)
	}
	//if we could read it ,try to convert its value to an int
	target, _ = strconv.Atoi(cookie.Value)

	guessVal, _ := strconv.Atoi(yourGuess)
	//compare YourGuess to target guess(random number)
	if guessVal== target{
		count=0
		winner=true
		message ="Correct Guess "+ yourGuess+" was the answer"
		cookie = &http.Cookie{
			Name: "target",
			Value: strconv.Itoa(rand.Intn(20-1)),
			Expires: time.Now().Add(72 * time.Hour),
		}	
		//set the cookie
		http.SetCookie(w,cookie)
	
	}else if guessVal < target {
	   message="Try Again your guess  was  too low"
		count += 1
		
	}else if guessVal > target {
		message="Try Again your guess was too high"
		count += 1	
	}

	cookie2 = &http.Cookie{
		Name: "count",
		Value: strconv.Itoa(count),
		Expires: time.Now().Add(72 * time.Hour),
	}
	http.SetCookie(w,cookie2)

	//read the contents of guess.html and return a template
	t, _ := template.ParseFiles("guess.tmpl")

	//execute template and pass pointer to myMsg 	struct
	t.Execute(w, &myMsg{Message:message,YourGuess:yourGuess,Winner:winner,Congrats:congrats})
}//guessHandler

func main() {
	// handles root page
	http.HandleFunc("/", requestHandler)

	//handle /guess page
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}
