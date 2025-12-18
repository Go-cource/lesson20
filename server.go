package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Task struct {
	Id           int
	CreationTime string
	Text         string
	Result       bool
	ResultTime   string
}

var Tasks = []Task{
	{Id: 1, CreationTime: time.Now().Format("2006/01/02 15:04:05"), Text: "Отжимания", Result: false, ResultTime: ""},
	{Id: 2, CreationTime: time.Now().Format("2006/01/02 15:04:05"), Text: "Подтягивания", Result: false, ResultTime: ""},
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("task/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, Tasks)
}
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("task/tasks.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, Tasks)
}
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("task/createTask.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, nil)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/create_task", createTaskHandler)
	http.ListenAndServe(":8081", nil)
}
