package dice

import (
	"fmt"
	"git.okki.hu/garric/roll/rng"
)

// Roll the Dice and return the result.
func (d Dice) Roll() int {
	return roll(int(d))
}

// RollMany rolls the Dice multiple times and sums up the results.
func (d Dice) RollMany(amount int) (result int) {
	for i := 0; i < amount; i++ {
		result += d.Roll()
	}
	return
}

// RollSpecial is the same as RollMany, but after the rolls
// a modifier is also added to the result.
func (d Dice) RollSpecial(amount, modifier int) int {
	return d.RollMany(amount) + modifier
}

// Roll the dice in given rpg dice definition (for example 3d6+2).
// Returns the result of the roll or an error of the def cannot be parsed.
func Roll(def string) (int, error) {
	count, sides, mod, err := parse(def)
	if err != nil {
		return 0, err
	}
	d, ok := SupportedDice()[sides]
	if !ok {
		return 0, fmt.Errorf("dice with %d sides are not supported", sides)
	}
	return d.RollSpecial(count, mod), nil
}

func roll(sides int) int {
	return rng.RandProvider.Intn(1, sides+1)
}
