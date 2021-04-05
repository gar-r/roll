package dice

import (
	"math/rand"
	"testing"
)

func Test_Dice_Roll(t *testing.T) {

	t.Run("d4 tests", func(t *testing.T) {
		assertRoll(t, &Dice {1, 4, 0}, 2)
		assertRoll(t, &Dice {2, 4, 0}, 6)
		assertRoll(t, &Dice {3, 4, 0}, 10)
		assertRoll(t, &Dice {4, 4, 0}, 14)
		assertRoll(t, &Dice {5, 4, 0}, 16)
	})

	t.Run("d6 tests", func(t *testing.T) {
		assertRoll(t, &Dice {1, 6, 0}, 6)
		assertRoll(t, &Dice {2, 6, 0}, 10)
		assertRoll(t, &Dice {3, 6, 0}, 16)
		assertRoll(t, &Dice {4, 6, 0}, 22)
		assertRoll(t, &Dice {5, 6, 0}, 24)
	})

	t.Run("d10 tests", func(t *testing.T) {
		assertRoll(t, &Dice {1, 10, 0}, 2)
		assertRoll(t, &Dice {2, 10, 0}, 10)
		assertRoll(t, &Dice {3, 10, 0}, 18)
		assertRoll(t, &Dice {4, 10, 0}, 28)
		assertRoll(t, &Dice {5, 10, 0}, 30)
	})

	t.Run("edge cases", func(t *testing.T) {
		assertRoll(t, &Dice {0, 4, 0}, 0)
		assertRoll(t, &Dice {0, 6, 0}, 0)
	})
}

func assertRoll(t *testing.T, d *Dice, expected int) {
	t.Helper()
	r := testRand()		// fixed seed for testing
	actual := d.Roll(r)
	if actual != expected {
		t.Errorf("expected %d, got %d. Rolls: %v", expected, actual, getRolls(d))
	}
}

func getRolls(d *Dice) []int {
	r := testRand()
	result := make([]int, 0)
	for i:=0; i<d.count; i++ {
		result = append(result, d.rollOne(r))
	}
	return result
}

func testRand() *rand.Rand {
	return rand.New(rand.NewSource(1))
}
