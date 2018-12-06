package main

import (
	"math"
	"strconv"
	"strings"
)

const TASK2_THRESHOLD = 10000

type Point struct {
	X             int
	Y             int
	ClosestToEdge bool
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func parseLine(line string) *Point {
	ps := strings.Split(line, ", ")
	x, err := strconv.Atoi(ps[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(ps[1])
	if err != nil {
		panic(err)
	}
	return NewPoint(x, y)
}

func parseInput(in chan string) (ps map[*Point]int, w, h int) {
	ps = make(map[*Point]int)
	for line := range in {
		p := parseLine(line)
		ps[p] = 0
		if p.X > w {
			w = p.X
		}
		if p.Y > h {
			h = p.Y
		}
	}
	w++
	h++
	return
}

func distance(a, b *Point) int {
	ax := float64(a.X)
	bx := float64(b.X)
	ay := float64(a.Y)
	by := float64(b.Y)
	return int(math.Abs(ax-bx) + math.Abs(ay-by))
}

func task1(in chan string) string {
	ps, w, h := parseInput(in)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			minD := math.MaxInt32
			var cp *Point
			foundOne := true
			for p, _ := range ps {
				d := distance(NewPoint(x, y), p)
				if d < minD {
					minD = d
					cp = p
					foundOne = true
				} else if d == minD {
					foundOne = false
					cp = nil
				}
			}
			if cp == nil {
				continue
			}
			if x == 0 || y == 0 || x == w-1 || y == h-1 {
				cp.ClosestToEdge = true
			}
			if foundOne {
				ps[cp]++
			}
		}
	}

	maxA := 0
	for p, a := range ps {
		if p.ClosestToEdge {
			continue
		}
		if a > maxA {
			maxA = a
		}
	}

	return strconv.Itoa(maxA)
}

func task2(in chan string) string {
	ps, w, h := parseInput(in)
	count := 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			totalD := 0
			for p, _ := range ps {
				totalD += distance(NewPoint(x, y), p)
			}
			if totalD < TASK2_THRESHOLD {
				count++
			}
		}
	}
	return strconv.Itoa(count)
}
