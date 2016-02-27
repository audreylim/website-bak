package main

import (
	"html/template"
	"net/http"
	"os"
)

var t = template.Must(template.New("index").ParseFiles("static/html/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", Index)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
