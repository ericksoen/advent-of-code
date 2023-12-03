package main

import (
	"testing"
)

func TestFinder(t *testing.T) {

	scenarios := []struct {
		token              string
		expected           []int
		useWordReplacement bool
	}{
		{"1a1", []int{1, 1}, false},
		{"oneight", []int{1, 8}, true},
		{"1sdf7ew2", []int{1, 7, 2}, false},
		{"sdf7", []int{7}, false},
		{"bpcfzztwo252", []int{2, 2, 5, 2}, true},
		{"one", []int{1}, true},
		{"eightwo", []int{8, 2}, true},
		{"dssmtmrkonedbbhdhjbf9hq", []int{1, 9}, true},
		{"dssmtmrkonedbbhdhjbf9hq", []int{9}, false},
		{"one41tvgttqnm1791szxcjbg", []int{4, 1, 1, 7, 9, 1}, false},
	}

	for _, scenario := range scenarios {

		searchItems := numericSearchItems

		if scenario.useWordReplacement {
			searchItems = append(searchItems, alphaSearchItems...)
		}
		actualValue := finder(scenario.token, searchItems)

		assertCorrectSlice(t, scenario.expected, actualValue)
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
