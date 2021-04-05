package main

import (
	"fmt"
	"testing"
)

func Test_Dice_Parse(t *testing.T) {

	// valid dice defs
	defs := []struct {
		def          string
		count, sides int
	}{
		// standard D&D dice
		{"d4", 1, 4},
		{"d6", 1, 6},
		{"d8", 1, 8},
		{"d10", 1, 10},
		{"d12", 1, 12},
		{"d20", 1, 20},
		{"d100", 1, 100},

		// single die, with count specified
		{"1d4", 1, 4},
		{"1d6", 1, 6},
		{"1d8", 1, 8},
		{"1d10", 1, 10},
		{"1d12", 1, 12},
		{"1d20", 1, 20},
		{"1d100", 1, 100},

		// multiple dice
		{"2d4", 2, 4},
		{"3d6", 3, 6},
		{"4d8", 4, 8},
		{"5d10", 5, 10},
		{"6d12", 6, 12},
		{"7d20", 7, 20},
		{"8d100", 8, 100},
	}

	for _, test := range defs {
		t.Run(test.def, func(t *testing.T) {
			assertDice(t, test.def, test.count, test.sides)
		})
	}

	// invalid dice defs: should return error when parsed
	invalid := []string{
		"", "1d", "1dx", "xd4",	// non-numbers
		"foo", 					// no "d"
		"2d3", "3d7", "4d11",	// no such die
		"1d14", "1d16", "1d18", // could be legal, but not
	}

	for _, test := range invalid {
		t.Run(fmt.Sprintf("def '%s'", test), func(t *testing.T) {
			assertError(t, test)
		})
	}

}

func assertError(t *testing.T, def string) {
	t.Helper()
	c, s, err := Parse(def)
	if err == nil {
		t.Errorf("expected error, got (%d, %d)", c, s)
	}
}

func assertDice(t *testing.T, def string, count, sides int) {
	t.Helper()
	c, s, err := Parse(def)
	if err != nil {
		t.Error(err)
	} else if count != c {
		t.Errorf("expected count: %d, got %d", count, c)
	} else if sides != s {
		t.Errorf("expected sides: %d, got %d", sides, s)
	}
}
