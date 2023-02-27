package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := struct {
		Title string
		Body  string
	}{
		Title: "User Profile",
		Body:  "User ID: " + id,
	}
	tmpl := template.Must(template.ParseFiles("templates/user.html"))
	tmpl.Execute(w, data)
}
