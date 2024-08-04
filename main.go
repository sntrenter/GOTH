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

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	//render the settings page using the template settings in the file hompage.templ
	response := settings().Render
	fmt.Fprint(w, response)
	log.Println("settings was clicked")
}

func main() {
	//component := hello("John")
	name := "John"
	age := 21
	isAdult := age >= 18

	component := homepage(name, age, isAdult)

	//settings := settings()

	http.Handle("/", templ.Handler(component))
	http.HandleFunc("/button-click", buttonClickHandler)
	http.HandleFunc("/Settings", SettingsHandler)

	fmt.Println("Listening on :3000")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
