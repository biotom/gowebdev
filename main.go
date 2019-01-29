//package gowebdev
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

_ "src/github.com/mattn/go-sqlite3"
)
type Page struct{
	Name string
	DBStatus bool
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html" ))
	db, err := sql.Open("sqlite3", "dev.db")
		if err != nil{
			fmt.Println(err)
		}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name:"Gopher"}
		if name := r.FormValue("name"); name != ""{
			p.Name = name
		}

		p.DBStatus = db.Ping() == nil

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8081", nil))
}