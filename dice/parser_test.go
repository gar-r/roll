package dice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {

	validDice := []struct {
		def                    string
		count, sides, modifier int
	}{
		{"d4", 1, 4, 0},
		{"d6", 1, 6, 0},
		{"d8", 1, 8, 0},
		{"d10", 1, 10, 0},
		{"d12", 1, 12, 0},
		{"d20", 1, 20, 0},
		{"d100", 1, 100, 0},

		{"D4", 1, 4, 0},
		{"D6", 1, 6, 0},
		{"D8", 1, 8, 0},
		{"D10", 1, 10, 0},
		{"D12", 1, 12, 0},
		{"D20", 1, 20, 0},
		{"D100", 1, 100, 0},

		{"1d4", 1, 4, 0},
		{"1d6", 1, 6, 0},
		{"1d8", 1, 8, 0},
		{"1d10", 1, 10, 0},
		{"1d12", 1, 12, 0},
		{"1d20", 1, 20, 0},
		{"1d100", 1, 100, 0},

		{"1D4", 1, 4, 0},
		{"1D6", 1, 6, 0},
		{"1D8", 1, 8, 0},
		{"1D10", 1, 10, 0},
		{"1D12", 1, 12, 0},
		{"1D20", 1, 20, 0},
		{"1D100", 1, 100, 0},

		{"2d4", 2, 4, 0},
		{"3d6", 3, 6, 0},
		{"4d8", 4, 8, 0},
		{"5d10", 5, 10, 0},
		{"6d12", 6, 12, 0},
		{"7d20", 7, 20, 0},
		{"8d100", 8, 100, 0},

		{"1d4+2", 1, 4, 2},
		{"1d6+3", 1, 6, 3},
		{"1d8+4", 1, 8, 4},
		{"1d10+5", 1, 10, 5},
		{"1d12+6", 1, 12, 6},
		{"1d20+10", 1, 20, 10},
		{"1d100+100", 1, 100, 100},

		{"1d4-5", 1, 4, -5},
		{"1d6-25", 1, 6, -25},
		{"1d8-100", 1, 8, -100},
	}

	for _, test := range validDice {
		t.Run(test.def, func(t *testing.T) {
			c, s, m, err := parse(test.def)
			assert.NoError(t, err)
			assert.Equal(t, test.count, c)
			assert.Equal(t, test.sides, s)
			assert.Equal(t, test.modifier, m)
		})
	}

	invalid := []string{
		"", "1d", "1dx", "xd4", // non-numbers
		"foo",                // no "d"
		"2d3", "3d7", "4d11", // no such die
		"1d14", "1d16", "1d18", // could be legal, but not
	}

	for _, test := range invalid {
		t.Run(fmt.Sprintf("def '%s'", test), func(t *testing.T) {
			_, _, _, err := parse(test)
			assert.Error(t, err)
		})
	}
}
