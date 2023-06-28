package main

import (
	"fmt"
	"git.okki.hu/garric/roll/rng"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateRolls(t *testing.T) {

	rng.RandProvider = &mockRandomProvider{}

	t.Run("default test cases", func(t *testing.T) {

		tests := []testCase{
			{[]string{"1d4", "1d6"}, []int{1, 1}},
			{[]string{"2d10", "1d8", "1d6"}, []int{2, 1, 1}},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
				assertRolls(t, test)
			})
		}
	})

	t.Run("negFlag disabled", func(t *testing.T) {
		negFlag = false
		test := testCase{[]string{"1d4-10"}, []int{0}}
		assertRolls(t, test)
	})

	t.Run("negFlag enabled", func(t *testing.T) {
		negFlag = true
		test := testCase{[]string{"1d4-10"}, []int{-9}}
		assertRolls(t, test)
	})
}

func assertRolls(t *testing.T, test testCase) {
	t.Helper()
	actual, err := generateRolls(test.input)
	assert.NoError(t, err)
	assert.Equal(t, test.expected, actual)
}

type testCase struct {
	input    []string
	expected []int
}

type mockRandomProvider struct{}

func (m *mockRandomProvider) Intn(_, _ int) int {
	return 1
}
