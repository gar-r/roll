package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_GenerateRolls(t *testing.T) {

	t.Run("default test cases", func(t *testing.T) {

		tests := []testCase{
			{[]string{"1d4", "1d6"}, []int{2, 4}},
			{[]string{"2d10", "1d8", "1d6"}, []int{10, 8, 6}},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
				assertRolls(t, test)
			})
		}
	})

	t.Run("negFlag disabled", func(t *testing.T) {
		negFlag = false
		test := testCase{ []string{"1d4-10"}, []int{0}}
		assertRolls(t, test)
	})

	t.Run("negFlag enabled", func(t *testing.T) {
		negFlag = true
		test := testCase{ []string{"1d4-10"}, []int{-8}}
		assertRolls(t, test)
	})
}

func assertRolls(t *testing.T, test testCase) {
	t.Helper()
	r = rand.New(rand.NewSource(1))
	actual, err := generateRolls(test.input)
	if err != nil {
		t.Error(err)
	}
	if len(actual) != len(test.expected) {
		t.Errorf("size mismatch: %v, %v", actual, test.expected)
	}
	for i, item := range actual {
		if test.expected[i] != item {
			t.Errorf("expected %d, got %d (%v, %v)", test.expected[i], item, test.expected, actual)
		}
	}
}

type testCase struct {
	input    []string
	expected []int
}
