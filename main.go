package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "idk what you're lookin' for ¯\\_(ツ)_/¯")
	}
}

func init() {
	log.Println("initiated ...")
	// static assets
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// handlers
	http.HandleFunc("/news", NewsHandler)
	http.HandleFunc("/", HomeHandler)
}

func main() {

}
