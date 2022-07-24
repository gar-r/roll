package main

import (
	"flag"
	"math/rand"
	"time"

	"roll/dice"
)

var sumFlag bool
var negFlag bool

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	initFlags()
	if rolls, err := generateRolls(flag.Args()); err != nil {
		exitErr(err)
	} else {
		if sumFlag {
			printSum(rolls)
		} else {
			printRolls(rolls)
		}
	}
	return
}

func generateRolls(args []string) ([]int, error) {
	rolls := make([]int, 0)
	for _, s := range args {
		if d, err := dice.Parse(s); err != nil {
			return nil, err
		} else {
			roll := d.Roll(r)
			if roll < 0 && !negFlag {
				roll = 0
			}
			rolls = append(rolls, roll)
		}
	}
	return rolls, nil
}
