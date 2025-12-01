package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_enterFloat(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{"Вводим число 10", "10", 10, false},
		{"Вводим дробь 10.5", "10.5", 10.5, false},
		{"Вводим отрицательное число -10", "-10", -10, false},
		{"Вводим буквы ab", "ab", 0, true},
		{"Ничего не вводим", "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdin = strings.NewReader(tt.input + "\n")
			got, gotErr := enterFloat(tt.name)

			if tt.wantErr && gotErr == nil {
				t.Errorf("enterFloat() failed: %v", tt.input)
			}
			if !tt.wantErr && gotErr != nil {
				t.Errorf("enterFloat() failed: %v", gotErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("enterFloat() failed: %v != %v", got, tt.want)
			}

		})
	}
}

func Test_divide(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{"Делим положительные 6 на 2", 6, 2, 3, false},
		{"Делим на отрицательное 10 на -5", 10, -5, -2, false},
		{"Делим отрицательные -100 на -10", -100, -10, 10, false},
		{"Делим 15 на 2.5", 15, 2.5, 6, false},
		{"Делим c дробным остатком 10 на 4", 10, 4, 2.5, false},
		{"Делим 10 на 0", 10, 0, 0, true},
		{"Делим 0 на 0", 0, 0, 0, true},
		{"Делим 0 на 10", 0, 10, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := divide(tt.a, tt.b)

			if tt.wantErr {
				if gotErr == nil {
					t.Errorf("divide() failed: %v", gotErr)
				}
			}
			if !tt.wantErr {
				if gotErr != nil {
					t.Errorf("divide() failed: %v", gotErr)
				}
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("divide() failed: %v != %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name    string
		values  string
		want    float64
		wantErr bool
	}{
		{"a=6 b=2", "6\n2\n", 3.00, false},
		{"a=10 =-5", "10\n-5\n", -2.00, false},
		{"a=-100 b=-10", "-100\n-10\n", 10.00, false},
		{"a=15 b=2.5", "15\n2.5\n", 6.00, false},
		{"a=10 b=4", "10\n4\n", 2.50, false},
		{"a=10 b=0", "10\n0\n", 0.00, true},
		{"a=0 b=0", "0\n0\n", 0.00, true},
		{"a=0 b=10", "0\n10\n", 0.00, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdin = strings.NewReader(tt.values)
			var output bytes.Buffer
			stdout = &output

			main()
			gotOutput := output.String()

			if tt.wantErr {
				if !strings.Contains(gotOutput, "Ошибка:") {
					t.Errorf("main() failed: %v не содержит 'Ошибка:'", gotOutput)
				}
			} else {
				if strings.Contains(gotOutput, "Ошибка:") {
					t.Errorf("main() failed: %v содержит 'Ошибка:'", gotOutput)
				}
				if strings.Contains(gotOutput, fmt.Sprintf("%f", tt.want)) {
					t.Errorf("main() failed: %s, ожидался %.2f", gotOutput, tt.want)
				}
			}
		})
	}
}
