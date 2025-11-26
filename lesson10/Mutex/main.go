/*
Будут генерироваться ip адреса:
Будет Шареный ресурс - map[string]int - ключ - ip, значение - количество обращений
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
	var SecureMap sync.Map - потокобезопасная мапа

	v, err SecureMap.Load(key)
	SecureMap.Set(key, value)
	SecureMap.Delete(key)
*/

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitSite(ip string) {
	v.mu.Lock()
	count := v.visited[ip]
	count++
	v.visited[ip] = count
	v.mu.Unlock()
}
func main() {
	// var mu sync.Mutex
	site := Visited{visited: make(map[string]int)}
	for i := 0; i < 5; i++ {
		go site.VisitSite(fmt.Sprintf("%v.%[1]v.%[1]v.%[1]v", rand.Intn(255)))
	}
	// fmt.Println(site)
}
