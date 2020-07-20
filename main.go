package main

import (
	"fmt"
	"html/template"
	_ "log"
	"net/http"
	// Sent Value u need to import template
)

// Product is fucked
// ! when u create struct u need to Comment over them
type Product struct {
	Name  string
	Price int
}

func main() {
	// route
	http.HandleFunc("/", index)
	// http.HandleFunc("/index", index)

	http.HandleFunc("/home", home)

	http.HandleFunc("/login", login)
	http.HandleFunc("/save", save)


	// runsever with PORT
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	//! run to index.html careful about path
	//! cant use with ParseFiles
	// http.ServeFile(w, r, "index.html")


	// sent data
	var templates = template.Must(template.ParseFiles("index.html"))
	myProduct := Product{"Milk", 40}
	templates.ExecuteTemplate(w, "index.html", myProduct)

}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "login.html")
}

func save(w http.ResponseWriter, r *http.Request) {
	// receive Request with Request
	fmt.Println("Method: ", r.Method)
	r.ParseForm()
	fmt.Println("Username: " , r.Form["firstname"])
	fmt.Println("Password: ", r.Form["lastname"])

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " this is Home page my friend hey")
}



// use gin for auto reload
// gin "gin -a 8000 -p 8080 run (filename).go"
// -a = app port / -p = web app port