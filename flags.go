package main

import (
	"flag"
	"fmt"
	"os"
)

func initFlags() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "\troll [roll-def] ...")
		fmt.Fprintln(os.Stderr, "\t\troll-def: [count]d<sides>[+modifier]")
		fmt.Fprintln(os.Stderr, "\t\tsides: [4,6,8,10,12,20,100]")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintln(os.Stderr, "\td6\trolls one six-sided die")
		fmt.Fprintln(os.Stderr, "\t1d10\trolls one ten-sided die")
		fmt.Fprintln(os.Stderr, "\t1d10+1\trolls one ten-sided die and adds 1 to the result")
		fmt.Fprintln(os.Stderr, "\t1d10-1\trolls one ten-sided die and subtracts 1 from the result")
		fmt.Fprintln(os.Stderr, "\t3d4\trolls three four-sided dice")
		fmt.Fprintln(os.Stderr, "\td4 d6\trolls a four-sided die and a six-sided die")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
	}
	flag.BoolVar(&sumFlag, "s", false, "print a single line summing up all dice")
	flag.BoolVar(&negFlag, "n", false, "allow rolls to go below 0 when using negative modifiers")
	flag.Parse()
}
