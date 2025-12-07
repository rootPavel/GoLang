package main

import (
	"fmt"
	"net/http"
)

var (
	token = "12345"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Write([]byte("Succes POST request"))
	} else if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"status": "Succes request"}`))
		w.Header()
	}

}

func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
