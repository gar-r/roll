package main

import (
	"flag"
	"math/rand"
	"time"
)

var sumFlag bool

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
		if count, sides, err := Parse(s); err != nil {
			return nil, err
		} else {
			rolls = append(rolls, Roll(r, count, sides))
		}
	}
	return rolls, nil
}
