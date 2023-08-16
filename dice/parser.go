package dice

import (
	"fmt"
	"regexp"
	"strconv"
)

var diceRe = regexp.MustCompile(`^(\d*)[D|d](100|20|12|10|8|6|4)([+|\-]\d+)?$`)

func parse(def string) (count, sides, mod int, err error) {
	match := diceRe.FindStringSubmatch(def)
	if len(match) == 0 {
		err = newError(def)
		return
	}
	count, err = parseCount(match[1], def)
	if err != nil {
		return
	}
	sides, err = parseSides(match[2], def)
	if err != nil {
		return
	}
	mod, err = parseModifier(match[3], def)
	return
}

func parseCount(s, def string) (count int, err error) {
	if s == "" {
		count = 1
		return
	}
	if count, err = strconv.Atoi(s); err != nil {
		err = newError(def)
	}
	return
}

func parseSides(s, def string) (sides int, err error) {
	if sides, err = strconv.Atoi(s); err != nil {
		err = newError(def)
	}
	return
}

func parseModifier(s, def string) (mod int, err error) {
	if s == "" {
		mod = 0
		return
	}
	if mod, err = strconv.Atoi(s); err != nil {
		err = newError(def)
	}
	return
}

func newError(def string) error {
	return fmt.Errorf("cannot parse dice definition: '%s'", def)
}
