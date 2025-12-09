package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var (
	token = 12345
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, token)

}

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("Token")
	if err != nil {
		http.Error(w, "There are no cookies", http.StatusUnauthorized)
		return
	}
	if cookie.Value != strconv.Itoa(token) {
		fmt.Fprintf(w, "Ваш Token=%v не подходит для авторизации", cookie.Value)
	}
	if cookie.Value == strconv.Itoa(token) {
		fmt.Fprintf(w, "Secret\nВаш Token=%v подходит для авторизации", cookie.Value)
	}
}

func main() {
	http.HandleFunc("/auth", AuthHandler)
	http.HandleFunc("/secret", SecretHandler)
	fmt.Println("Server start ...")
	http.ListenAndServe(":8080", nil)
}
