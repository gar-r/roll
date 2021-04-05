package main

import (
	"errors"
	"regexp"
	"strconv"
)

var diceRe = regexp.MustCompile(`^(\d*)d(100|20|12|10|8|6|4)$`)

func Parse(def string) (count, sides int, err error) {
	match := diceRe.FindStringSubmatch(def)
	if match == nil || len(match) == 0 {
		err = errors.New("invalid def: " + def)
		return
	}
	if match[1] == "" {
		count = 1
	} else {
		if count, err = strconv.Atoi(match[1]); err != nil {
			return
		}
	}
	sides, err = strconv.Atoi(match[2])
	return
}
