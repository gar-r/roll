package dice

import "math/rand"

type Dice struct {
	count int
	sides int
	modifier int
}

func Parse(def string) (*Dice, error) {
	count, sides, modifier, err := parseDice(def)
	if err != nil {
		return nil, err
	}
	return &Dice { count, sides, modifier }, nil
}

func (d *Dice) Roll(r *rand.Rand) int {
	sum := 0
	for i:=0; i<d.count; i++ {
		sum += d.rollOne(r)
	}
	return sum
}

func (d *Dice) rollOne(r *rand.Rand) int {
	return r.Intn(d.sides) + 1
}
