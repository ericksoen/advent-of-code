package main

import (
	"testing"
)

func TestFinder(t *testing.T) {

	scenarios := map[string][]int{
		"1sdf7ew2":     {1, 7, 2},
		"sdf7":         {7},
		"bpcfzztwo252": {2, 5, 2},
		// "one":      {1},
	}

	for token, expectedValue := range scenarios {
		actualValue := finder(token)

		assertCorrectSlice(t, expectedValue, actualValue)
	}

}

func assertCorrectSlice(t testing.TB, expected, got []int) {
	t.Helper()

	if len(expected) != len(got) {
		t.Errorf("Slice lengths do not match. Expected %d. Got %d.", len(expected), len(got))
	}

	for i, value := range expected {

		if value != got[i] {
			t.Errorf("Error at index position = %d. Expected %+v. Got %+v", i, expected, got)
		}
	}
}
