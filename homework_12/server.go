package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "modernc.org/sqlite"
)

type Accounts struct {
	Id       int
	Email    string
	Password string
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./index.html")
		if err != nil {
			fmt.Println(err)
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		email := r.PostFormValue("email")
		pass := r.PostFormValue("password")
		fmt.Println(email, pass)

		connect, err := sql.Open("sqlite", "./accounts.db")
		if err != nil {
			fmt.Println(err)
		}
		defer connect.Close()

		rows, err := connect.Query("SELECT * FROM users WHERE email = ?", email)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		Acc := []Accounts{}
		for rows.Next() {
			a := Accounts{}
			err := rows.Scan(&a.Id, &a.Email, &a.Password)
			if err != nil {
				fmt.Println(err)
			}
			Acc = append(Acc, a)
		}

		if len(Acc) == 0 || Acc[0].Email != email {
			w.Write([]byte("Пользователь не найден"))
			return
		} else if Acc[0].Email == email {
			if Acc[0].Password != pass {
				w.Write([]byte("Пароль неверный"))
			} else if Acc[0].Password == pass {
				w.Write([]byte("Пароль верный"))
			}
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth", AuthHandler)
	http.ListenAndServe("0.0.0.0:8080", mux)
}
