package main

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(def string) (count, sides int, err error) {
	if empty(def) {
		err = errors.New("empty dice def")
		return
	}

	idx := strings.Index(def, "d")
	if idx < 0 {
		err = errors.New("missing 'd' separator")
		return
	}

	// parse dice count
	seg := def[:idx]
	if empty(seg) {
		count = 1
	} else {
		if count, err = strconv.Atoi(seg); err != nil {
			return
		}
	}

	// parse dice sides
	seg = def[idx+1:]
	if empty(seg) {
		err = errors.New("missing side count")
		return
	}
	sides, err = strconv.Atoi(def[idx+1:])
	if sides % 2 != 0 {
		err = errors.New("invalid die def: " + def)
	}
	return
}

func empty(def string) bool {
	return "" == def
}
