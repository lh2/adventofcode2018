package main

import (
	"strconv"
	"unicode"
)

func reactPolymer(polymer string) (result string, reactions int) {
	runes := []rune(polymer)
	reaction := false
	for i := 0; i < len(runes)-1; i++ {
		a := runes[i]
		b := runes[i+1]
		reaction = a != b &&
			unicode.ToLower(a) == unicode.ToLower(b)
		if reaction {
			reactions++
			i++
		} else {
			result += string(polymer[i])
		}

	}
	if !reaction {
		result += string(runes[len(runes)-1])
	}
	return
}

func task1(in chan string) string {
	polymer := <-in
	reactions := 1
	for reactions > 0 {
		polymer, reactions = reactPolymer(polymer)
	}
	return strconv.Itoa(len(polymer))
}
