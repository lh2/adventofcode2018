package main

import (
	"fmt"
	"strconv"
)

const SEARCH_PATTERN_TRESHOLD = 1000
const PATTERN_SAMPLE_SIZE = 200

type Cell int

const (
	CellEmpty Cell = iota
	CellTree
	CellLumber
)

type Grid struct {
	data    []Cell
	dataBuf []Cell
	Width   int
	Height  int
}

func (g *Grid) At(x, y int) Cell {
	return g.data[y*g.Width+x]
}

func (g *Grid) Swap() {
	old := g.data
	g.data = g.dataBuf
	g.dataBuf = old
}

func (g *Grid) Counts() (ne, nt, nl int) {
	for _, c := range g.data {
		switch c {
		case CellEmpty:
			ne++
		case CellTree:
			nt++
		case CellLumber:
			nl++
		}
	}
	return
}

func (g *Grid) Print() {
	for i, c := range g.data {
		if i%g.Height == 0 {
			fmt.Printf("\n")
		}
		r := '.'
		switch c {
		case CellTree:
			r = '|'
		case CellLumber:
			r = '#'
		}
		fmt.Printf(string(r))
	}
	fmt.Printf("\n")
}

func (g *Grid) NCounts(x, y int) (ne, nt, nl int) {
	for py := y - 1; py <= y+1; py++ {
		for px := x - 1; px <= x+1; px++ {
			if px < 0 || py < 0 ||
				px == g.Width || py == g.Height ||
				(px == x && py == y) {
				continue
			}
			switch g.At(px, py) {
			case CellEmpty:
				ne++
			case CellTree:
				nt++
			case CellLumber:
				nl++
			}
		}
	}

	return
}

func (g *Grid) Tick() {
	for i, c := range g.data {
		y := i / g.Height
		x := i % g.Height
		_, nt, nl := g.NCounts(x, y)
		switch c {
		case CellEmpty:
			if nt >= 3 {
				c = CellTree
			}
		case CellTree:
			if nl >= 3 {
				c = CellLumber
			}
		case CellLumber:
			if nl == 0 || nt == 0 {
				c = CellEmpty
			}
		}
		g.dataBuf[i] = c
	}
	g.Swap()
}

func parseInput(in chan string) *Grid {
	data := make([]Cell, 0)
	width := 0
	height := 0
	for line := range in {
		for i, r := range []rune(line) {
			var c Cell
			switch r {
			case '|':
				c = CellTree
			case '#':
				c = CellLumber
			}
			data = append(data, c)
			if i > width {
				width = i
			}
		}
		height++
	}
	width++
	return &Grid{
		data:    data,
		dataBuf: make([]Cell, len(data)),
		Width:   width,
		Height:  height,
	}
}

func task1(in chan string) string {
	g := parseInput(in)
	//g.Print()
	for i := 0; i < 10; i++ {
		g.Tick()
		//g.Print()
	}
	_, nt, nl := g.Counts()
	return strconv.Itoa(nt * nl)
}

func task2(in chan string) string {
	g := parseInput(in)
	samples := make([]int, 0)
	i := 0
	for ; i < SEARCH_PATTERN_TRESHOLD + PATTERN_SAMPLE_SIZE + 1; i++ {
		g.Tick()
		if i >= SEARCH_PATTERN_TRESHOLD {
			_, nt, nl := g.Counts()
			rv := nt * nl
			samples = append(samples, rv)
		}
	}
	seed := samples[0]
	var pl int
	for i, sv := range samples[1:] {
		if (sv == seed) {
			pl = i + 1
			break
		}
	}

	offset := (1000000000 - SEARCH_PATTERN_TRESHOLD) % pl
	return strconv.Itoa(samples[pl+offset-1])
}
