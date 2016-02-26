package main

import (
	"html/template"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index").ParseFiles("static/html/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
//}

func main() {
	http.HandleFunc("/", Index)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	//	router := httprouter.New()
	//	router.GET("/", Index)
	//	//	router.GET("/hello/:name", Hello)
	//	router.ServeFiles("/static/*filepath", http.Dir("static"))
	//
	//	log.Fatal(http.ListenAndServe(":8080", router))
}
