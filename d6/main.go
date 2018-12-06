package main

import (
	"math"
	"strings"
	"strconv"
)

type Point struct {
	X int
	Y int
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

func distance(a, b *Point) int {
	ax := float64(a.X)
	bx := float64(b.X)
	ay := float64(a.Y)
	by := float64(b.Y)
	return int(math.Abs(ax - bx) + math.Abs(ay - by))
}

func task1(in chan string) string {
	ps := make(map[*Point]int)
	w := 0
	h := 0
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
	