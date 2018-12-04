package main

import (
	"regexp"
	"strconv"
)

const CLOTH_SIZE = 1000
var cutRegexp = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

func mustAtoi(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}

func task1(in chan string) string {
	cloth := make([][]int, CLOTH_SIZE)
	for i := 0; i < CLOTH_SIZE; i++ {
		cloth[i] = make([]int, CLOTH_SIZE)
	}
	for line := range in {
		grps := cutRegexp.FindStringSubmatch(line)
		//id := mustAtoi(grps[1])
		x := mustAtoi(grps[2])
		y := mustAtoi(grps[3])
		w := mustAtoi(grps[4])
		h := mustAtoi(grps[5])
		for xp := x; xp < x + w; xp++ {
			for yp := y; yp < y + h; yp++ {
				cloth[xp][yp]++
			}
		}
	}
	cntInches := 0
	for _, row := range cloth {
		for _, cell := range row {
			if cell > 1 {
				cntInches++
			}
		}
	}
	return strconv.Itoa(cntInches)
}