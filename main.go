package main

import (
	"flag"

	"github.com/gar-r/roll/dice"
)

var (
	sumFlag bool
	negFlag bool
)

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
}

func generateRolls(args []string) ([]int, error) {
	rolls := make([]int, 0)
	for _, s := range args {
		roll, err := dice.Roll(s)
		if err != nil {
			return nil, err
		}
		if roll < 0 && !negFlag {
			roll = 0
		}
		rolls = append(rolls, roll)
	}
	return rolls, nil
}
