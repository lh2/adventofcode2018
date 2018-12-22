package main

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type RegionType int

const (
	RegionTypeRocky RegionType = iota
	RegionTypeWet
	RegionTypeNarrow
)

type Cave struct {
	Depth  int
	Target Point

	gidxCache map[Point]int
	elvlCache map[Point]int
}

func (c Cave) GeoIndex(p Point) (idx int) {
	if idx, ok := c.gidxCache[p]; ok {
		return idx
	}
	switch {
	case p == Point{} || p == c.Target:
		idx = 0
	case p.Y == 0:
		idx = p.X * 16807
	case p.X == 0:
		idx = p.Y * 48271
	default:
		idx = c.ErosionLevel(pt(p.X-1, p.Y)) *
			c.ErosionLevel(pt(p.X, p.Y-1))
	}
	c.gidxCache[p] = idx
	return
}

func (c Cave) ErosionLevel(p Point) (lvl int) {
	if lvl, ok := c.elvlCache[p]; ok {
		return lvl
	}
	lvl = (c.GeoIndex(p) + c.Depth) % 20183
	c.elvlCache[p] = lvl
	return
}

func (c Cave) RegionType(p Point) RegionType {
	return RegionType(c.ErosionLevel(p) % 3)
}

func pt(x, y int) Point {
	return Point{X: x, Y: y}
}

func parseInput(in chan string) (cave Cave, target Point) {
	cave = Cave{
		Depth:     mustAtoi((<-in)[7:]),
		gidxCache: make(map[Point]int),
		elvlCache: make(map[Point]int),
	}
	tc := strings.Split((<-in)[8:], ",")
	target = Point{
		X: mustAtoi(tc[0]),
		Y: mustAtoi(tc[1]),
	}
	cave.Target = target
	return
}

func task1(in chan string) string {
	c, target := parseInput(in)
	riskLevel := 0
	for x := 0; x <= target.X; x++ {
		for y := 0; y <= target.Y; y++ {
			riskLevel += int(c.RegionType(pt(x, y)))
		}
	}
	return strconv.Itoa(riskLevel)
}
