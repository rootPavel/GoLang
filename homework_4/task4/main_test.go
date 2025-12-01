package main

import "testing"

func Test_sumAll(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{"Сумма 1,2,3", []int{1, 2, 3}, 6},
		{"Сумма 10,-2,4,7", []int{10, -2, 4, 7}, 19},
		{"Сумма одного числа", []int{5}, 5},
		{"Сумма отрицательных чисел", []int{-1, -2, -3}, -6},
		{"Сумма пустого списка", []int{}, 0},
		{"Сумма больших чисел", []int{10000000, 2000000, 300000}, 12300000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumAll(tt.numbers...)
			if got != tt.want {
				t.Errorf("sumAll() failed: sum %v != %v want", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	main()
}
