package main

import (
	"math/rand"
	"testing"
)

func Test_Roll(t *testing.T) {

	t.Run("d4 tests", func(t *testing.T) {
		assertRoll(t, 1, 4, 2)
		assertRoll(t, 2, 4, 6)
		assertRoll(t, 3, 4, 10)
		assertRoll(t, 4, 4, 14)
		assertRoll(t, 5, 4, 16)
	})

	t.Run("d6 tests", func(t *testing.T) {
		assertRoll(t, 1, 6, 6)
		assertRoll(t, 2, 6, 10)
		assertRoll(t, 3, 6, 16)
		assertRoll(t, 4, 6, 22)
		assertRoll(t, 5, 6, 24)
	})

	t.Run("d10 tests", func(t *testing.T) {
		assertRoll(t, 1, 10, 2)
		assertRoll(t, 2, 10, 10)
		assertRoll(t, 3, 10, 18)
		assertRoll(t, 4, 10, 28)
		assertRoll(t, 5, 10, 30)
	})

	t.Run("edge cases", func(t *testing.T) {
		assertRoll(t, 0, 4, 0)
		assertRoll(t, 0, 6, 0)
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
