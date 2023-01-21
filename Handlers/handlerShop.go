package Handlers

import (
	"html/template"
	"net/http"
)

func Shop(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/shop.html"))
	tmpl.Execute(w, getUser(w))
}
