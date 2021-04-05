package main

import (
	"math/rand"
	"testing"
)

func Test_Roll(t *testing.T) {

	t.Run("d4 tests", func(t *testing.T) {
		assertRoll(t, 1, 4, 3)
		assertRoll(t, 2, 4, 4)
	})

	t.Run("edge cases", func(t *testing.T) {
	})
}

func assertRoll(t *testing.T, count, sides, expected int) {
	t.Helper()
	r := testRand()		// fixed seed for testing
	actual := Roll(r, count, sides)
	if actual != expected {
		t.Errorf("expected %d, got %d. Rolls: %v", expected, actual, getRolls(count, sides))
	}
}

func getRolls(count, sides int) []int {
	r := testRand()
	result := make([]int, 0)
	for i:=0; i<count; i++ {
		result = append(result, rollOne(r, sides))
	}
	return result
}

func testRand() *rand.Rand {
	return rand.New(rand.NewSource(1))
}
