package main

import (
	"errors"
	"strconv"
	"strings"
)

type Dice struct {
	count, sides int
}

func Parse(def string) (*Dice, error) {
	if empty(def) {
		return nil, errors.New("empty dice def")
	}
	parts := strings.Split(def, "d")
	count := 1
	if parts[0] != "" {
		c, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		count = c
	}
	sides, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	return &Dice{count, sides}, nil
}

func empty(def string) bool {
	return "" == def
}
