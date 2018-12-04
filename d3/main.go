package main

import (
	"regexp"
	"strconv"
	"fmt"
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

func parseCut(line string) (id, x, y, w, h int) {
	grps := cutRegexp.FindStringSubmatch(line)
	id = mustAtoi(grps[1])
	x = mustAtoi(grps[2])
	y = mustAtoi(grps[3])
	w = mustAtoi(grps[4])
	h = mustAtoi(grps[5])
	return
}

func computeClothOverlap(list []string) [][]int {
	cloth := make([][]int, CLOTH_SIZE)
	for i := 0; i < CLOTH_SIZE; i++ {
		cloth[i] = make([]int, CLOTH_SIZE)
	}
	for _, line := range list {
		_, x, y, w, h := parseCut(line)
		for xp := x; xp < x + w; xp++ {
			for yp := y; yp < y + h; yp++ {
				cloth[xp][yp]++
			}
		}
	}
	return cloth
}

func task1(in chan string) string {
	cloth := computeClothOverlap(inAsSlice(in))
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

func task2(in chan string) string {
	list := inAsSlice(in)
	cloth := computeClothOverlap(list)
	uniqueId := -1
	for _, line := range list {
		id, x, y, w, h := parseCut(line)
		hasOverlap := false
		for xp := x; !hasOverlap && xp < x + w; xp++ {
			for yp := y; !hasOverlap && yp < y + h; yp++ {
				hasOverlap = cloth[xp][yp] > 1
			}
		}
		if !hasOverlap {
			uniqueId = id
			break
		}
	}
	return strconv.Itoa(uniqueId)
}