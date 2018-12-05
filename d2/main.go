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

func diffIDs(id1, id2 string) (int, string) {
	diff := 0
	sameChars := make([]rune, 0)
	for i, c1 := range id1 {
		c2 := []rune(id2)[i]
		if c1 != c2 {
			diff++
		} else {
			sameChars = append(sameChars, c1)
		}
	}
	return diff, string(sameChars)
}

func task2(in chan string) string {
	list := inAsSlice(in)
	lowestDiff := int(^uint(0) >> 1)
	var sameChars string
	for k1, l1 := range list {
		for _, l2 := range list[k1:] {
			var diff int
			diff, chars := diffIDs(l1, l2)
			if diff < lowestDiff {
				lowestDiff = diff
				sameChars = chars
			}
		}
	}
	return sameChars
}
