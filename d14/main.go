package main

import (
	"strconv"
	"strings"
)

const INIT_SCORES = 37

func toDigits(n int) []int {
	dgs := make([]int, 0)
	for _, d := range strings.Split(strconv.Itoa(n), "") {
		dgs = append(dgs, mustAtoi(d))
	}
	return dgs
}

func join(scores []int) (str string) {
	for _, score := range scores {
		str += strconv.Itoa(score)
	}
	return
}

func brew(scores, elfs []int) []int {
	sum := 0
	for _, elf := range elfs {
		sum += scores[elf]
	}
	scores = append(scores, toDigits(sum)...)
	for i, elf := range elfs {
		elfs[i] = (elf + 1 + scores[elf]) % len(scores)
	}
	return scores
}

func task1(in chan string) string {
	scores := toDigits(INIT_SCORES)
	rc := mustAtoi(<-in)
	elfs := []int{0, 1}
	for len(scores) <= rc+10 {
		scores = brew(scores, elfs)
	}
	return join(scores[rc : rc+10])
}

func compareScores(scores, search []int) (found, shifted bool) {
	scl := len(scores)
	sel := len(search)
	if scl-1 < sel {
		return
	}
	sc := scores[scl-sel-1:]
	found = true
	for i := 0; i < sel; i++ {
		if shifted {
			if sc[i+1] != search[i] {
				found = false
				return
			}
		} else if sc[i] != search[i] {
			shifted = true
			i = -1
		}
	}
	return
}

func task2(in chan string) string {
	scores := toDigits(INIT_SCORES)
	search := toDigits(mustAtoi(<-in))
	elfs := []int{0, 1}
	shifted := false
	for {
		scores = brew(scores, elfs)
		var found bool
		if found, shifted = compareScores(scores, search); found {
			break
		}
	}
	foundI := len(scores) - len(search) - 1
	if shifted {
		foundI -= 1
	}
	return strconv.Itoa(foundI)
}
