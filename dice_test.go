package main

import "testing"

func Test_Dice_Parse(t *testing.T) {

	// invalid dice defs: should return error when parsed
	invalid := []struct{ def, desc string }{
		{"", "empty string"},
	}

	for _, test := range invalid {
		t.Run(test.desc, func(t *testing.T) {
			assertError(t, test.def)
		})
	}

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

		// dice count
		{ "1d6", 1, 6 },
	}

	for _, test := range defs {
		t.Run(test.def, func(t *testing.T) {
			assertDice(t, test.def, test.count, test.sides)
		})
	}

}

func assertError(t *testing.T, def string) {
	t.Helper()
	_, err := Parse(def)
	if err == nil {
		t.Error("expected error, got none")
	}
}

func assertDice(t *testing.T, def string, count, sides int) {
	t.Helper()
	d, err := Parse(def)
	if err != nil {
		t.Error(err)
	} else if d == nil {
		t.Error("nil dice")
	} else if count != d.count {
		t.Errorf("expected count: %d, got %d", count, d.count)
	} else if sides != d.sides {
		t.Errorf("expected sides: %d, got %d", sides, d.sides)
	}
}
