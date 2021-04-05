package main

import "math/rand"

func Roll(r *rand.Rand, count, sides int) int {
	sum := 0
	for i:=0; i<count; i++ {
		sum += rollOne(r, sides)
	}
	return sum
}

func rollOne(r *rand.Rand, sides int) int {
	return r.Intn(sides) + 1
}
