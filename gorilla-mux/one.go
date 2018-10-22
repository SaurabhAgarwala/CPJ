package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

type Destination struct {
	Title string
	Visited  bool
}

type DestinationPageData struct {
	PageTitle string
	Destinations []Destination
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, you are currently in the following endpoint: %s\n", r.URL.Path)
		fmt.Fprintf(w, "Welcome to the basic implementaion of a web server using Golang.\n\n")
		fmt.Fprintf(w, "Lets know you, after locally serving the app, go to the URL localhost/detail/{name}/{age}\n")
		fmt.Fprintf(w, "To see the assets present go to the URL localhost/detail/{name}/{age}\n")
		fmt.Fprintf(w, "To see the basic implementation of a template go to the URL localhost/destinations")
	})

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
    r.PathPrefix("/static/").Handler(s)
    http.Handle("/", r)
	
	r.HandleFunc("/detail/{name}/{age}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := vars["age"]

		fmt.Fprintf(w, "You are %s and your age is %s\n", name, age)
	})

	tmpl := template.Must(template.ParseFiles("layout.html"))

	r.HandleFunc("/destinations", func(w http.ResponseWriter, r *http.Request) {
		data := DestinationPageData{
			PageTitle: "Destinations list",
			Destinations: []Destination{
				{Title: "Mumbai", Visited: false},
				{Title: "Delhi", Visited: true},
				{Title: "Bangalore", Visited: true},
			},
		}
		tmpl.Execute(w, data)
	})
	
	http.ListenAndServe(":80", r)
}