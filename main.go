package main

import (
	"fmt"
	// Sent Value u need to import template
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//! ResultData is fucked
// ! when u create struct u need to Comment over them
type ResultData struct {
	Id    int
	Age   int
	Fname string
	Lname string
}

// ! Careful about name's func you have to use Pascal case not camel case and not snake case neither

func main() {

	// route
	http.HandleFunc("/", index)

	http.HandleFunc("/home", home)

	http.HandleFunc("/login", login)
	http.HandleFunc("/save", save)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/update", update)

	// runsever with PORT
	http.ListenAndServe(":8000", nil)

}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
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
		results = append(results, tRes)
		if err != nil {
			panic(err)
		}
	}

	// sent data
	var templates = template.Must(template.ParseFiles("index.html"))
	templates.Execute(w, results)
	// fmt.Println(results)
}

func delete(res http.ResponseWriter, req *http.Request) {
	var db, err = sql.Open("mysql", "root:@/golang")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// for delelet
	stmt, err := db.Prepare("delete from users where id=?")
	stmt.Exec(req.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}
	log.Println("Deleted successfully")
	defer db.Close()
	http.Redirect(res, req, "/", 301)
}

func edit(res http.ResponseWriter, req *http.Request) {
	db := dbConn()
	uId := req.URL.Query().Get("id")
	slDB, err := db.Query("SELECT * from users WHERE id=?", uId)
	if err != nil {
		panic(err.Error())
	}
	user := ResultData{}
	for slDB.Next() {
		var id, age int
		var fname, lname string
		err = slDB.Scan(&id, &fname, &lname, &age)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Fname = fname
		user.Lname = lname
		user.Age = age
	}
	var templates = template.Must(template.ParseFiles("edit.html"))
	templates.Execute(res, user)
	defer db.Close()
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func save(w http.ResponseWriter, r *http.Request) {
	// receive Request with Request r.Method = Request.Method
	// fmt.Println("Method: ", r.Method)
	// r.ParseForm >> Receive value from Form  / r.Form = Request.Form
	// r.ParseForm()
	// var Fname = r.Form["firstname"]
	// var Lname = r.Form["lastname"]
	// var age = r.Form["age"]

	// fmt.Println(Fname, Lname, age)
	// // connect to DB
	// var db, err = sql.Open("mysql", "root:@/golang")
	// if err != nil {
	// 	panic(err)
	// }
	// // defer db.Close()
	// smtm, err := db.Prepare("insert into users (firstname , lastname , age) values (?,?,?)")
	// smtm.Exec(Fname, Lname, age)
	// if err != nil {
	// 	panic(err)
	// }
	// http.Redirect(w, r, "/", 301)
	db := dbConn()
	if r.Method == "POST" {
		Fname := r.FormValue("firstname")
		Lname := r.FormValue("lastname")
		Age := r.FormValue("age")
		insForm, err := db.Prepare("Insert Into users (firstname,lastname,age) values(?,?,?)")
			if err != nil {
				panic(err.Error())
			}
		insForm.Exec(Fname, Lname, Age)
		log.Println("Insert | Name : " + Fname + "| Lame : " + Lname + "| age : " + Age)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func update(res http.ResponseWriter, req *http.Request) {
	db := dbConn()
		if req.Method == "POST" {
			ID := req.FormValue("uid")
			Fname := req.FormValue("firstname")
			Lname := req.FormValue("lastname")
			Age := req.FormValue("age")
			uPDB, err := db.Prepare("UPDATE USERS SET firstname=? , lastname=? , age=? WHERE id=?")
				if err != nil {
					panic(err.Error())
				}
			uPDB.Exec(Fname, Lname, Age, ID)
			log.Println("Update | Name : " + Fname + "| Lastname : " + Lname + "| age : " + Age )
		}
	defer db.Close()
	http.Redirect(res, req, "/", 301)
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
