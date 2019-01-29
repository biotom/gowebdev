//package gowebdev
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	_ "src/github.com/mattn/go-sqlite3"
)

type SearchResult struct{
	Title string
	Author string
	Year string
	ID string
}
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

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request){
		results := []SearchResult{
			SearchResult{"Moby-Dick", "Herman Welville", "1851", "2222" },
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w,err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8081", nil))
}