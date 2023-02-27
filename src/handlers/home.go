package handlers

import (
	"html/template"
	"net/http"
)

type User struct {
	Id int
	Name string
	Age  int
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		Body  string
		Users []User
	}{
		Title: "My Website",
		Body:  "Welcome to my website. Powered by GO!",
		Users : []User {
			{Id: 1, Name: "Alice", Age: 25},
			{Id: 2, Name: "Bob", Age: 30},
			{Id: 3, Name: "Charlie", Age: 35},
		},
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}
