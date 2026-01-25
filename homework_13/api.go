package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Data struct {
	Data []CoinCap `json:"data"`
}

type CoinCap struct {
	Name     string `json:"name"`
	PriceUsd string `json:"priceUsd"`
}

func main() {
	fmt.Println("Введите API keys для доступа к https://rest.coincap.io/v3/asset:")
	stringOut := bufio.NewReader(os.Stdin)
	rowApiKey, err := stringOut.ReadString('\n')
	if err != nil {
		log.Println("Not read string", err)
		return
	}

	apiKey := strings.TrimSpace(rowApiKey)
	if apiKey == "" {
		fmt.Println("Вы не ввели Api Key попробуйте еще раз!")
		return
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, "https://rest.coincap.io/v3/assets", nil)
	if err != nil {
		log.Println("Error create requst: ", err)
		return
	}

	setBearer := fmt.Sprintf("Bearer %s", apiKey)
	request.Header.Set("Authorization", setBearer)

	response, err := client.Do(request)
	if err != nil {
		log.Println("Error do request: ", err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error read body: ", err)
	}

	if response.StatusCode != 200 {
		log.Printf("Status code = %v \n%s\n", response.StatusCode, string(body))
	}

	var price Data

	if err := json.Unmarshal(body, &price); err != nil {
		log.Println("Error unmarshal ", err)
	}

	for _, usd := range price.Data {
		fmt.Printf("Name: %s\nPrice: %s\n...................................\n", usd.Name, usd.PriceUsd)
	}

}
