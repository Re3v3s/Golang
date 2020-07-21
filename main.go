package main

import (
	"fmt"
	// Sent Value u need to import template
	"html/template"
	_ "log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//! ResultData is fucked
// ! when u create struct u need to Comment over them
type ResultData struct {
	 Id int
	 Age int
	 Fname string
	 Lname string
}

func main() {

	// route
	http.HandleFunc("/", index)

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

	//! connect
	var db, err = sql.Open("mysql", "root:@/golang")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//! query SQL
	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println(err)
	}

	tRes := ResultData{}
	var results []ResultData
	//! loop with Next()
	for rows.Next() {
		var Id int
		var Age int
		var Fname string
		var Lname string

		err = rows.Scan(&Id, &Fname, &Lname, &Age)
		tRes.Id = Id
		tRes.Fname = Fname
		tRes.Lname = Lname
		tRes.Age = Age
		results = append(results,tRes)
			if err != nil {
				panic(err)
			}
	}

	// sent data
	var templates = template.Must(template.ParseFiles("index.html"))
	templates.Execute(w,results)
	fmt.Println(results)
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func save(w http.ResponseWriter, r *http.Request) {
	// receive Request with Request
	fmt.Println("Method: ", r.Method)
	r.ParseForm()
	fmt.Println("Username: ", r.Form["firstname"])
	fmt.Println("Password: ", r.Form["lastname"])

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " this is Home page my friend hey")
}

// ! for auto reload
// use gin for auto reload
// gin "gin -a 8000 -p 8080 run (filename).go"
// -a = app port / -p = web app port

// ! for SQL
//  go get -u github.com/go-sql-driver/mysql
