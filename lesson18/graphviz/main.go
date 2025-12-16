package main

import (
	"net/http"
	_ "net/http/pprof"
)

const (
	maxSize = 1_000_000
)

func foo() {
	for {
		var s []int
		// s := make([]int, 0, maxSize)
		for i := 0; i < maxSize; i++ {
			s = append(s, i)
		}
	}
}
func main() {
	go foo()
	http.ListenAndServe(":8080", nil)
}

// go tool pprof -http=":9090" -seconds=30 http://localhost:8080/debug/pprof/profile
// go tool pprof -http=":9090" -seconds=30 http://localhost:8080/debug/pprof/heap

// curl -v http://localhost:8080/debug/pprof/heap > heap.out
// go tool pprof -http=":9090" -seconds=30 heap.out
