package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}
type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO LIST CHANGE",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Make the demo", Done: false},
		},
	}
	tmpl.Execute(w, data)
}

func handerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)
	log.Fatal(http.ListenAndServe(":9000", mux))
}
