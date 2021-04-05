package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		if _, _, err := Parse(arg); err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	}
}
