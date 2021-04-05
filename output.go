package main

import (
	"fmt"
	"os"
)

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

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
