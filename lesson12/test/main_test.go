package main

import "testing"

func TestIsEven(t *testing.T) {
	t.Parallel()
	result := IsEven(1)
	if result != "no" {
		t.Errorf("Error in test - expexted: 'no', got: %s", result)
	}
	t.Log("Test1 finished")

	resultTrue := IsEven(2)
	if resultTrue != "yes" {
		t.Errorf("Error in test - expexted: 'yes', got: %s", resultTrue)
	}
	t.Log("Test2 finished")
}

func TestIsEven2(t *testing.T) {
	t.Parallel()
	var testcases = []struct {
		text string
		got  int
		want string
	}{
		{"-5 - отрицательное", -5, "no"},
		{"0 - zero", 0, "yes"},
		{"2 - Even", 2, "yes"},
		{"1 - Not Even", 1, "no"},
		{"500123456 - Not Even", 500123456, "yes"},
	}
	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := IsEven(tc.got)
			if result != tc.want {
				t.Errorf("Error in test %s: got: %s want: %s", tc.text, result, tc.want)
			}
		})
	}
}
