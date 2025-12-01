package main

import "testing"

func BenchmarkInsertXInMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(100)
	}
}

func BenchmarkInsertXInMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(1000)
	}
}
func BenchmarkInsertXInMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(100000)
	}
}

// func BenchmarkInsertXInMapInterface100(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		InsertXInMap(100)
// 	}
// }

// func BenchmarkInsertXInMapInterface1000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		InsertXInMap(1000)
// 	}
// }
// func BenchmarkInsertXInMapInterface100000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		InsertXInMap(100000)
// 	}
// }

func BenchmarkInsertXInSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(100)
	}
}

func BenchmarkInsertXInSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(1000)
	}
}
func BenchmarkInsertXInSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(100000)
	}
}
