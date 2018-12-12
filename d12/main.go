package main

import (
	"strconv"
)

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

func run(rounds int, pots []int, rules map[int]int) (sum int) {
	for i := 0; i < rounds; i++ {
		pots = applyRules(rules, hashedPots(pots))
	}
	for i, pot := range pots {
		i -= rounds * 2
		if pot == 1 {
			sum += i
		}
	}
	return
}

func task1(in chan string) string {
	pots, rules := parseInput(in)
	sum := run(20, pots, rules)
	return strconv.Itoa(sum)
}
