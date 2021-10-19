package main

import "testing"

type TestCase struct {
	Description string
	Input       int
	Expected    int
}

func TestDouble(t *testing.T) {

	testCases := []TestCase{
		{"Doubling 3", 3, 6},
		{"Doubling 4", 4, 8},
		{"Doubling 5", 5, 10},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			result := double(testCase.Input)
			if result != testCase.Expected {
				t.Errorf("Expected %d got %d", testCase.Expected, result)
			}
		})
	}
}

/*
func TestDouble(t *testing.T) {
	result := double(3)

	if result != 6 {
		t.Errorf("Expected %d got %d", 6, result)
	}
}
*/
