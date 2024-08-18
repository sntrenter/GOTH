package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func buttonClickHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here, e.g., increment a counter or any other business logic
	response := "Button was clicked!"
	fmt.Fprint(w, response)
	log.Println("button was clicked")
}

func tableHandler(w http.ResponseWriter, r *http.Request) {
	//response := table()
	//fmt.Fprint(w, response.Render())
	//log.Println("table was clicked")
	test := getRowsAsHTML(0)
	log.Println("table was clicked")
	fmt.Fprint(w, test)
	//fmt.Fprint(w, "<p>fake table data</p>")
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	//render the settings page using the template settings in the file hompage.templ
	//response := settings()
	//fmt.Fprint(w, response.Render())
	fmt.Fprint(w, templ.Handler(settings()))
	log.Println("settings was clicked")
}

func addRowHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, addRow(r.FormValue("url")))
}

func main() {

	component := homepage()
	//settings := settings()

	http.Handle("/", templ.Handler(component))
	http.Handle("/Settings", templ.Handler(settings()))
	http.HandleFunc("/button-click", buttonClickHandler)
	http.HandleFunc("/tableHandler", tableHandler)
	http.HandleFunc("/addRow", addRowHandler)

	fmt.Println("Listening on :3000")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
