// На вход программе (в функции main) передается список из ip адрсов []string{"192.168.1.1, 192.168.1.2, 192.168.1.1, 192.168.1.3, 192.168.1.4"}
// 		Нужно вывести на экран:
// 			1) Сколько всего было обращений
// 			2) Сколько уникальных обращений
// 			3) Кто самый популярный ip адрес

package main

import (
	"fmt"
)

func main() {
	listIp := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.1",
		"192.168.1.3",
		"192.168.1.4",
		"192.168.1.3",
		"192.168.1.4",
		"192.168.1.1",
	}

	uniqueRequest := CountIpRequest(listIp)
	popularIp := PopularIp(uniqueRequest)

	fmt.Printf("Всего было обращений: %d\n", len(listIp))
	fmt.Printf("Уникальных обращений: %d\n", len(uniqueRequest))
	fmt.Printf("Cамый популярный ip адрес: %s, обращений: %d \n", popularIp, uniqueRequest[popularIp])
}

func CountIpRequest(listIp []string) map[string]int {
	uniqueRequest := make(map[string]int)
	for _, ip := range listIp {
		uniqueRequest[ip]++
	}
	return uniqueRequest
}

func PopularIp(mapIp map[string]int) string {
	count := 0
	var popularIp string
	for ip, request := range mapIp {
		if request > count {
			count = request
			popularIp = ip
		}
	}
	return popularIp
}
