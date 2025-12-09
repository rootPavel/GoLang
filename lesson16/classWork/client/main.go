package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	request_auth, err := http.Get("http://localhost:8080/auth")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer request_auth.Body.Close()

	body_auth, err := io.ReadAll(request_auth.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	token := string(body_auth)
	fmt.Printf("Получили токен - %v\n", token)

	request_secret, err := http.NewRequest("GET", "http://localhost:8080/secret", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_secret.Header.Add("Cookie", "Token="+token)
	client := &http.Client{}
	response_secret, err := client.Do(request_secret)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response_secret.Body.Close()
	body_secret, err := io.ReadAll(response_secret.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body_secret))

}
