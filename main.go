package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	type Contacts struct {
		Email   string
		Phone   string
		Address string
	}
	OurContacts := Contacts{
		Email:   "veryovkin@bmstu.ru",
		Phone:   "+7 (999) 123-32-12",
		Address: "Москва, ул. Радио д. 1",
	}
	tmpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, OurContacts)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error: ", err)
	}
}
