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
	s1 := def[:idx]
	if empty(s1) {
		count = 1
	} else {
		if count, err = strconv.Atoi(s1); err != nil {
			return
		}
	}

	// parse dice sides
	s2 := def[idx+1:]
	if empty(s2) {
		err = errors.New("missing side count")
		return
	}
	sides, err = strconv.Atoi(def[idx+1:])
	return
}

func empty(def string) bool {
	return "" == def
}
