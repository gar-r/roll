package dice

import (
	"fmt"
	"testing"
)

func Test_Dice_Parse(t *testing.T) {

	// valid dice defs
	defs := []struct {
		def                    string
		count, sides, modifier int
	}{
		// standard D&D dice
		{"d4", 1, 4, 0},
		{"d6", 1, 6, 0},
		{"d8", 1, 8, 0},
		{"d10", 1, 10, 0},
		{"d12", 1, 12, 0},
		{"d20", 1, 20, 0},
		{"d100", 1, 100, 0},

		// single die, with count specified
		{"1d4", 1, 4, 0},
		{"1d6", 1, 6, 0},
		{"1d8", 1, 8, 0},
		{"1d10", 1, 10, 0},
		{"1d12", 1, 12, 0},
		{"1d20", 1, 20, 0},
		{"1d100", 1, 100, 0},

		// multiple dice
		{"2d4", 2, 4, 0},
		{"3d6", 3, 6, 0},
		{"4d8", 4, 8, 0},
		{"5d10", 5, 10, 0},
		{"6d12", 6, 12, 0},
		{"7d20", 7, 20, 0},
		{"8d100", 8, 100, 0},
	}

	for _, test := range defs {
		t.Run(test.def, func(t *testing.T) {
			assertDice(t, test.def, test.count, test.sides, test.modifier)
		})
	}

	// invalid dice defs: should return error when parsed
	invalid := []string{
		"", "1d", "1dx", "xd4", 	// non-numbers
		"foo",                		// no "d"
		"2d3", "3d7", "4d11", 		// no such die
		"1d14", "1d16", "1d18", 	// could be legal, but not
	}

	for _, test := range invalid {
		t.Run(fmt.Sprintf("def '%s'", test), func(t *testing.T) {
			assertError(t, test)
		})
	}
}

func assertError(t *testing.T, def string) {
	t.Helper()
	c, s, m, err := parseDice(def)
	if err == nil {
		t.Errorf("expected error, got (%d, %d, %d)", c, s, m)
	}
}

func assertDice(t *testing.T, def string, count, sides, modifier int) {
	t.Helper()
	c, s, m, err := parseDice(def)
	if err != nil {
		t.Error(err)
	} else if count != c {
		t.Errorf("expected count: %d, got %d", count, c)
	} else if sides != s {
		t.Errorf("expected sides: %d, got %d", sides, s)
	} else if modifier != m {
		t.Errorf("expected modifier: %d, got %d", modifier, m)
	}
}
