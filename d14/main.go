package main

import (
	"strconv"
	"strings"
)

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

func task1(in chan string) string {
	scores := []int{3, 7}
	rc := mustAtoi(<-in)
	elfs := []int{0, 1}
	for len(scores) <= rc+10 {
		sum := 0
		for _, elf := range elfs {
			sum += scores[elf]
		}
		scores = append(scores, toDigits(sum)...)
		for i, elf := range elfs {
			elfs[i] = (elf + 1 + scores[elf]) % len(scores)
		}
	}
	return join(scores[rc : rc+10])
}
