package main

import (
	"math"
	"regexp"
	"strconv"
)

var lineRegex = regexp.MustCompile(`position=\<\s*([\d-]+),\s*([\d-]+)\> velocity=\<\s*([\d-]+),\s*([\d-]+)\>`)

type Point struct {
	X    int64
	Y    int64
	VelX int64
	VelY int64
}

func pointFromLine(line string) *Point {
	m := lineRegex.FindStringSubmatch(line)
	return &Point{
		X:    mustAtoi64(m[1]),
		Y:    mustAtoi64(m[2]),
		VelX: mustAtoi64(m[3]),
		VelY: mustAtoi64(m[4]),
	}
}

func pointsFromInput(in chan string) []*Point {
	ps := make([]*Point, 0)
	for line := range in {
		ps = append(ps, pointFromLine(line))
	}
	return ps
}

func computeArea(ps []*Point) (a, x, y, w, h int64) {
	x = math.MaxInt64
	y = math.MaxInt64
	for _, p := range ps {
		if p.X < x {
			x = p.X
		}
		if p.X > w {
			w = p.X
		}
		if p.Y < y {
			y = p.Y
		}
		if p.Y > h {
			h = p.Y
		}
	}
	a = (w - x) * (h - y)
	return
}

func computeTick(ps []*Point) []*Point {
	psn := make([]*Point, len(ps))
	for i, p := range ps {
		pn := *p
		pn.X += pn.VelX
		pn.Y += pn.VelY
		psn[i] = &pn
	}
	return psn
}

func draw(ps []*Point, xs, ys, w, h int64) (s string) {
	rows := make([][]bool, int(h-ys)+1)
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]bool, int(w-xs)+1)
	}
	for _, p := range ps {
		rows[p.Y-ys][p.X-xs] = true
	}
	for _, r := range rows {
		for _, c := range r {
			if c {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return
}

func searchMessage(ps []*Point) (msg string, i int) {
	la := int64(math.MaxInt64)
	var ops []*Point
	i = -1
	for {
		a, x, y, w, h := computeArea(ps)
		if a > la {
			msg = draw(ops, x, y, w, h)
			return
		}
		la = a
		ops = ps
		ps = computeTick(ps)
		i++
	}
}

func task1(in chan string) string {
	ps := pointsFromInput(in)
	msg, _ := searchMessage(ps)
	return msg
}

func task2(in chan string) string {
	ps := pointsFromInput(in)
	_, i := searchMessage(ps)
	return strconv.Itoa(i)
}
