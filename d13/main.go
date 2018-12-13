package main

import (
	"fmt"
	"sort"
)

type Direction int

const (
	DirLeft = iota
	DirUp
	DirRight
	DirDown

	DirMax
)

type Cart struct {
	X   int
	Y   int
	Dir Direction
	//DirUp is interpreted as "straight" here
	NextTurn Direction
}

func (c *Cart) Remove() {
	c.X = -1
	c.Y = -1
}

type Carts []*Cart

func (cs Carts) Len() int {
	return len(cs)
}

func (cs Carts) Less(i, j int) bool {
	a := cs[i]
	b := cs[j]
	if a.Y < b.Y {
		return true
	}
	if a.Y == b.Y && a.X < b.X {
		return true
	}
	return false
}

func (cs Carts) Swap(i, j int) {
	tmp := cs[i]
	cs[i] = cs[j]
	cs[j] = tmp
}

func (cs Carts) At(x, y int) *Cart {
	for _, c := range cs {
		if c.X == x && c.Y == y {
			return c
		}
	}
	return nil
}

type Crash struct {
	X int
	Y int
}

type Railway struct {
	Grid   [][]rune
	Carts  Carts
	Width  int
	Height int
}

func (rw *Railway) Print() {
	for y, row := range rw.Grid {
		for x, cell := range row {
			if c := rw.Carts.At(x, y); c != nil {
				fmt.Printf("X")
			} else {
				fmt.Printf(string(cell))
			}
		}
		fmt.Printf("\n")
	}
}

func (rw *Railway) Tick() []Crash {
	crashes := make([]Crash, 0)
	sort.Sort(rw.Carts)

	for _, cart := range rw.Carts {
		if cart.X < 0 || cart.Y < 0 {
			continue
		}
		x := cart.X
		y := cart.Y
		switch cart.Dir {
		case DirUp:
			y--
		case DirRight:
			x++
		case DirDown:
			y++
		case DirLeft:
			x--
		}
		if x < 0 || y < 0 || x >= rw.Width || y >= rw.Height {
			panic(fmt.Sprintf("Cart tries to escape into %d,%d, impossible", x, y))
		}
		oc := rw.Carts.At(x, y)
		if oc != nil {
			oc.Remove()
			cart.Remove()
			crashes = append(crashes, Crash{
				X: x,
				Y: y,
			})
			continue
		}
		d := cart.Dir
		switch rw.Grid[y][x] {
		case '/':
			switch d {
			case DirUp:
				d = DirRight
			case DirRight:
				d = DirUp
			case DirDown:
				d = DirLeft
			case DirLeft:
				d = DirDown
			}
		case '\\':
			switch d {
			case DirUp:
				d = DirLeft
			case DirRight:
				d = DirDown
			case DirDown:
				d = DirRight
			case DirLeft:
				d = DirUp
			}
		case '+':
			switch cart.NextTurn {
			case DirLeft:
				d = (DirMax + d - 1) % DirMax
			case DirRight:
				d = (DirMax + d + 1) % DirMax
			}
			cart.NextTurn = (cart.NextTurn + 1) % DirDown
		case '|':
			if d != DirUp && d != DirDown {
				panic(fmt.Sprintf("Cart crashed into rails at %d,%d", x, y))
			}
		case '-':
			if d != DirLeft && d != DirRight {
				panic(fmt.Sprintf("Cart crashed into rails at %d,%d", x, y))
			}
		case ' ':
			panic(fmt.Sprintf("Carts can only drive on rails at %d,%d", x, y))
		default:
			panic("This cannot happen.")
		}
		cart.Dir = d
		cart.X = x
		cart.Y = y
	}

	ncarts := make([]*Cart, 0, len(rw.Carts))
	for _, cart := range rw.Carts {
		if cart.X >= 0 && cart.Y >= 0 {
			ncarts = append(ncarts, cart)
		}
	}
	rw.Carts = ncarts

	return crashes
}

func parseInput(in chan string) *Railway {
	grid := make([][]rune, 0)
	carts := make([]*Cart, 0)
	var w int
	var h int
	for line := range in {
		row := []rune(line)
		if w == 0 {
			w = len(row)
		}
		for x, r := range row {
			var dir Direction
			var rail rune
			switch r {
			case '^':
				dir = DirUp
				rail = '|'
			case '>':
				dir = DirRight
				rail = '-'
			case 'v':
				dir = DirDown
				rail = '|'
			case '<':
				dir = DirLeft
				rail = '-'
			default:
				continue
			}
			carts = append(carts, &Cart{
				X:   x,
				Y:   h,
				Dir: dir,
			})
			row[x] = rail
		}
		grid = append(grid, row)
		h++
	}
	return &Railway{
		Grid:   grid,
		Carts:  carts,
		Width:  w,
		Height: h,
	}
}

func task1(in chan string) string {
	rw := parseInput(in)
	for {
		cs := rw.Tick()
		if len(cs) > 0 {
			return fmt.Sprintf("%d,%d\n", cs[0].X, cs[0].Y)
		}
	}
}

func task2(in chan string) string {
	rw := parseInput(in)
	for {
		_ = rw.Tick()
		if len(rw.Carts) == 1 {
			return fmt.Sprintf("%d,%d\n", rw.Carts[0].X, rw.Carts[0].Y)
		}
	}
}
