package main

import (
	"strconv"
)

func analyzeID(id string) map[rune]int {
	r := make(map[rune]int)
	for _, v := range id {
		r[v] = r[v] + 1
	}
	return r
}

func task1(in chan string) string {
	two := 0
	three := 0
	for line := range in {
		r := analyzeID(line)
		c2 := true
		c3 := true
		for _, v := range r {
			if !c2 && !c3 {
				break
			}
			if c2 && v == 2 {
				two++
				c2 = false
			}
			if c3 && v == 3 {
				three++
				c3 = false
			}
		}
	}
	return strconv.Itoa(two * three)
}