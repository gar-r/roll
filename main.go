package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
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

func printRolls(rolls []int) {
	for _, r := range rolls {
		fmt.Printf("%d\n", r)
	}
}

func printSum(rolls []int) {
	sum := 0
	for _, r := range rolls {
		sum += r
	}
	fmt.Printf("%d\n", sum)
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

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func initFlags() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "\troll <roll-def(s)>")
		fmt.Fprintln(os.Stderr, "\t\troll-def: [count]d<sides>")
		fmt.Fprintln(os.Stderr, "\t\tsides: [4,6,8,10,12,20,100]")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintln(os.Stderr, "\td6\trolls one six-sided die")
		fmt.Fprintln(os.Stderr, "\t1d10\trolls one ten-sided die")
		fmt.Fprintln(os.Stderr, "\t3d4\trolls three four-sided dice")
		fmt.Fprintln(os.Stderr, "\td4 d6\trolls a four-sided die and a six-sided die")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
	}
	flag.BoolVar(&sumFlag, "s", false, "print a single line summing up all dice")
	flag.Parse()
}
