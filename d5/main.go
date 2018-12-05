package main

import (
	"strconv"
	"strings"
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
			result += string(runes[i])
		}

	}
	if !reaction {
		result += string(runes[len(runes)-1])
	}
	return
}

func fullyReactPolymer(polymer string) string {
	reactions := 1
	for reactions > 0 {
		polymer, reactions = reactPolymer(polymer)
	}
	return polymer
}

func task1(in chan string) string {
	polymer := fullyReactPolymer(<-in)
	return strconv.Itoa(len(polymer))
}

func task2(in chan string) string {
	polymer := <-in
	shortest := int(^uint(0) >> 1)
	for unit := 'a'; unit <= 'z'; unit++ {
		r := strings.NewReplacer(
			string(unit), "",
			string(unicode.ToUpper(unit)), "")
		p := r.Replace(polymer)

		p = fullyReactPolymer(p)
		pl := len(p)
		if pl < shortest {
			shortest = pl
		}
	}
	return strconv.Itoa(shortest)
}
