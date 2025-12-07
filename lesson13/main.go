package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type MyStruct struct {
	Name  string
	Email string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprint(w, "404 page not found")
	}
	// http.ServeFile(w, r, "index.html")
	data := MyStruct{Name: "Alex", Email: "Alex@test.ru"} //fmt.Sprintf("%v", time.Now())
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, data)

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page")
}

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts page")
}

func main() {
	/*
		/ - index(main) page
		/about - about page
		/contacts - contacts
	*/
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contacts", ContactsHandler)
	fmt.Println("Server started...")
	http.ListenAndServe("localhost:8080", nil)
}
