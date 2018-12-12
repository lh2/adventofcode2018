package main

import (
	"strconv"
)

const PATTERN_FINDER_THRESHOLD = 10

func parseInput(in chan string) ([]int, map[int]int) {
	pots := make([]int, 0)
	ruleset := make(map[int]int)
	for _, r := range []rune((<-in)[15:]) {
		pot := 0
		if r == '#' {
			pot = 1
		}
		pots = append(pots, pot)
	}
	<-in
	for line := range in {
		var rule int
		shift := uint(4)
		for _, r := range []rune(line[:5]) {
			pot := 0
			if r == '#' {
				pot = 1
			}
			rule |= pot << shift
			shift--
		}
		resultPot := 0
		if line[9:] == "#" {
			resultPot = 1
		}
		ruleset[rule] = resultPot
	}
	return pots, ruleset
}

func hashedPots(pots []int) chan int {
	out := make(chan int)
	go func() {
		temp := [5]int{0, 0, 0, 0, 0}
		for _, pot := range pots {
			temp[0] |= pot << 0
			temp[1] |= pot << 1
			temp[2] |= pot << 2
			temp[3] |= pot << 3
			temp[4] |= pot << 4

			out <- temp[0]
			temp[0] = temp[1]
			temp[1] = temp[2]
			temp[2] = temp[3]
			temp[3] = temp[4]
			temp[4] = 0
		}
		out <- temp[0]
		out <- temp[1]
		out <- temp[2]
		out <- temp[3]
		close(out)
	}()
	return out
}

func applyRules(rules map[int]int, in chan int) []int {
	pots := make([]int, 0)
	for hash := range in {
		pots = append(pots, rules[hash])
	}
	return pots
}

func extractPattern(pots []int) []int {
	var s int
	for ; s < len(pots); s++ {
		if pots[s] > 0 {
			break
		}
	}
	e := len(pots) - 1
	for ; e > 0; e-- {
		if pots[e] > 0 {
			break
		}
	}
	return pots[s : e+1]
}

func comparePattern(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sum(rounds int, pots []int) (sum int) {
	for i, pot := range pots {
		i -= rounds * 2
		if pot == 1 {
			sum += i
		}
	}
	return
}

func run(rounds int, pots []int, rules map[int]int) int {
	var lastPattern []int
	var samePatternCount int
	var lastSum int
	var sumDiff int
	var i int
	for ; i < rounds; i++ {
		pots = applyRules(rules, hashedPots(pots))
		pattern := extractPattern(pots)
		if lastPattern != nil && comparePattern(lastPattern, pattern) {
			samePatternCount++
			sum := sum(i+1, pots)
			sumDiff = sum - lastSum
			lastSum = sum
		} else {
			samePatternCount = 0
		}
		if samePatternCount > PATTERN_FINDER_THRESHOLD {
			break
		}
		lastPattern = pattern
	}

	if samePatternCount == 0 {
		return sum(rounds, pots)
	}
	return (rounds-i-1)*sumDiff + lastSum
}

func task1(in chan string) string {
	pots, rules := parseInput(in)
	sum := run(20, pots, rules)
	return strconv.Itoa(sum)
}

func task2(in chan string) string {
	pots, rules := parseInput(in)
	sum := run(50000000000, pots, rules)
	return strconv.Itoa(sum)
}
