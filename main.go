package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error: ", err)
	}
}
