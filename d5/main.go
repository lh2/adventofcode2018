package main

import (
	"strconv"
	"strings"
	"unicode"
)

func reactPolymer(polymer string) string {
	runes := []rune(polymer)

	reaction := false
	for i := 0; i < len(runes)-1; i++ {
		a := runes[i]
		if a == 0 {
			continue
		}
		nextID := i
		var b rune
		for b == 0 {
			nextID++
			b = runes[nextID]
		}
		reaction = a != b &&
			unicode.ToLower(a) == unicode.ToLower(b)
		if reaction {
			runes[i] = 0
			runes[nextID] = 0
			ni := i
			for ; ni >= 0 && runes[ni] == 0; ni-- {}
			if ni < 0 {
				ni = i
				for ; runes[ni] == 0; ni++ {}
			}
			i = ni - 1
		}

	}
	
	result := make([]rune, 0, len(runes))
	for _, r := range runes {
		if r > 0 {
			result = append(result, r)
		}
	}
	
	return string(result)
}

func task1(in chan string) string {
	polymer := reactPolymer(<-in)
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

		p = reactPolymer(p)
		pl := len(p)
		if pl < shortest {
			shortest = pl
		}
	}
	return strconv.Itoa(shortest)
}
