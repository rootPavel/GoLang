package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type TasksStructure struct {
	Id           int
	Text         string
	CreationTime string
	Result       bool
	ResultTime   string
}

var (
	tasks = []TasksStructure{
		{Id: 1, Text: "Проснуться", CreationTime: "19.12.2025 08:07", Result: true, ResultTime: time.Now().Format("02.01.2006 15:04")},
		{Id: 2, Text: "Позавтракать", CreationTime: "19.12.2025 08:08", Result: true, ResultTime: time.Now().Format("02.01.2006 15:04")},
		{Id: 3, Text: "Сходить на работу", CreationTime: "19.12.2025 08:08", Result: true, ResultTime: time.Now().Format("02.01.2006 15:04")},
		{Id: 4, Text: "Поужинать", CreationTime: "19.12.2025 08:08", Result: false, ResultTime: ""},
		{Id: 5, Text: "Лечь спать", CreationTime: "19.12.2025 08:08", Result: false, ResultTime: ""},
	}
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Используйте метод GET", http.StatusMethodNotAllowed)
	}
	tmpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, tasks)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Используйте метод GET", http.StatusMethodNotAllowed)
	}
	tmpl, err := template.ParseFiles("pages/tasks.html")
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, tasks)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("pages/createTask.html")
		if err != nil {
			log.Println(err)
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		var newTask TasksStructure
		newTask.Id = len(tasks) + 1
		newTask.Text = r.FormValue("TaskText")
		newTask.CreationTime = time.Now().Format("02.01.2006 15:04")
		tasks = append(tasks, newTask)
		log.Println("Данные успешно записаны!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/tasks", tasksHandler)
	mux.HandleFunc("/create_task", createTaskHandler)

	log.Println("Start server...")
	http.ListenAndServe(":8080", mux)
}
