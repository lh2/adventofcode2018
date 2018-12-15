package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type UnitType int

const (
	UnitTypeElf UnitType = iota
	UnitTypeGoblin
)

type Point struct {
	X int
	Y int
}

func (p Point) Ns() [4]Point {
	return [4]Point{
		Point{p.X, p.Y - 1},
		Point{p.X - 1, p.Y},
		Point{p.X + 1, p.Y},
		Point{p.X, p.Y + 1},
	}
}

type Points []Point

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) Less(i, j int) bool {
	a := ps[i]
	b := ps[j]
	if a.Y < b.Y {
		return true
	}
	if a.Y == b.Y && a.X < b.X {
		return true
	}
	return false
}

func (ps Points) Swap(i, j int) {
	tmp := ps[i]
	ps[i] = ps[j]
	ps[j] = tmp
}

type Unit struct {
	Type UnitType
	AP   int
	HP   int
	X    int
	Y    int
}

func (u *Unit) Range() [4]Point {
	return Point{u.X, u.Y}.Ns()
}

type Units []*Unit

func (u Units) Len() int {
	return len(u)
}

func (u Units) Less(i, j int) bool {
	a := u[i]
	b := u[j]
	if a.Y < b.Y {
		return true
	}
	if a.Y == b.Y && a.X < b.X {
		return true
	}
	return false
}

func (u Units) Swap(i, j int) {
	tmp := u[i]
	u[i] = u[j]
	u[j] = tmp
}

type Cell struct {
	Wall bool
	Unit *Unit
}

type Game struct {
	Map   [][]*Cell
	Units Units
	mapw  int
	maph  int
}

func (g *Game) SearchEnemies(u *Unit) (es Units) {
	es = make(Units, 0)
	for y := 0; y < g.maph; y++ {
		for x := 0; x < g.mapw; x++ {
			e := g.Map[y][x].Unit
			if e == nil || e.Type == u.Type || e.HP < 0 {
				continue
			}
			es = append(es, e)
		}
	}
	return es
}

func (g *Game) AvailableTargets(us Units) Points {
	ps := make([]Point, 0)
	for _, u := range us {
		for _, p := range u.Range() {
			c := g.Map[p.Y][p.X]
			if !c.Wall && c.Unit == nil {
				ps = append(ps, p)
			}
		}
	}
	return ps
}

func (g *Game) CellFree(p Point) bool {
	return g.Map[p.Y][p.X].Unit == nil && !g.Map[p.Y][p.X].Wall
}

func (g *Game) genPathMap(pm [][]int, o Point, d int) {
	pm[o.Y][o.X] = d
	ins := o.Ns()
	nns := ins[:]
	for len(nns) > 0 {
		d++
		for _, p := range nns {
			if g.CellFree(p) {
				pm[p.Y][p.X] = d
			}
		}
		nnns := make(map[string]Point, 0)
		for _, pp := range nns {
			if !g.CellFree(pp) {
				continue
			}
			for _, p := range pp.Ns() {
				if pm[p.Y][p.X] < 0 && g.CellFree(p) {
					k := fmt.Sprintf("%d-%d", p.X, p.Y)
					if _, ok := nnns[k]; !ok {
						nnns[k] = p
					}
				}
			}
		}
		nns = make(Points, 0)
		for _, p := range nnns {
			nns = append(nns, p)
		}
	}
}

func (g *Game) printPathMap(pm [][]int) {
	panic("do not")
	for _, rows := range pm {
		for _, c := range rows {
			if c < 0 {
				fmt.Printf("   |")
			} else {
				fmt.Printf("%3d|", c)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (g *Game) MapPaths(o Point) [][]int {
	pm := make([][]int, 0)
	for y := 0; y < g.maph; y++ {
		row := make([]int, 0)
		for x := 0; x < g.mapw; x++ {
			row = append(row, -1)
		}
		pm = append(pm, row)
	}
	g.genPathMap(pm, o, 0)
	return pm
}

func (g *Game) PathToClosestTarget(o *Unit, ps Points) *Point {
	pm := g.MapPaths(Point{o.X, o.Y})
	cd := int(math.MaxInt32)
	var cps Points
	for _, p := range ps {
		d := pm[p.Y][p.X]
		if d < cd && d > -1 {
			cd = d
			cps = make(Points, 0)
		}
		if d == cd {
			cps = append(cps, p)
		}
	}
	sort.Sort(cps)
	if len(cps) == 0 {
		return nil
	}
	pm = g.MapPaths(cps[0])
	cd = int(math.MaxInt32)
	for _, p := range o.Range() {
		d := pm[p.Y][p.X]
		if d < cd && d > -1 {
			cd = d
			cps = make(Points, 0)
		}
		if d == cd {
			cps = append(cps, p)
		}
	}
	sort.Sort(cps)
	return &(cps[0])
}

func (g *Game) MoveUnit(u *Unit, np Point) {
	g.Map[u.Y][u.X].Unit = nil
	if g.Map[np.Y][np.X].Unit != nil || g.Map[np.Y][np.X].Wall {
		panic("Illegal Move")
	}
	g.Map[np.Y][np.X].Unit = u
	u.X = np.X
	u.Y = np.Y
}

func (g *Game) TargetInVicinity(u *Unit) *Unit {
	ts := make([]*Unit, 0)
	for _, p := range (Point{u.X, u.Y}).Ns() {
		c := g.Map[p.Y][p.X]
		if c.Wall || c.Unit == nil || c.Unit.Type == u.Type {
			continue
		}
		ts = append(ts, c.Unit)
	}
	if len(ts) == 0 {
		return nil
	}
	var nts Units
	lhp := int(math.MaxInt32)
	for _, u := range ts {
		if u.HP < lhp {
			lhp = u.HP
			nts = make([]*Unit, 0)
		}
		if lhp == u.HP {
			nts = append(nts, u)
		}
	}
	sort.Sort(nts)
	return nts[0]
}

func (g *Game) Attack(u, e *Unit) {
	e.HP -= u.AP
	if e.HP < 0 {
		g.Map[e.Y][e.X].Unit = nil
		e.Y = -1
		e.X = -1
	}
}

func (g *Game) Tick() bool {
	sort.Sort(g.Units)

	noop := true
	for _, u := range g.Units {
		if u.HP < 0 {
			continue
		}
		t := g.TargetInVicinity(u)
		if t != nil {
			g.Attack(u, t)
			noop = false
			continue
		}
		es := g.SearchEnemies(u)
		targets := g.AvailableTargets(es)
		np := g.PathToClosestTarget(u, targets)
		if np == nil {
			continue
		}
		noop = false
		if np.X == u.X && np.Y == u.Y {
			panic("Moving to same cell")
		}
		g.MoveUnit(u, *np)
		t = g.TargetInVicinity(u)
		if t != nil {
			g.Attack(u, t)
		}
	}
	return !noop
}

func (g *Game) Print() {
	for _, row := range g.Map {
		hpstr := "   "
		for _, c := range row {
			p := "."
			if c.Unit != nil {
				switch c.Unit.Type {
				case UnitTypeElf:
					p = "E"
				case UnitTypeGoblin:
					p = "G"
				}
				hpstr += fmt.Sprintf("%s(%d), ", p, c.Unit.HP)
			} else if c.Wall {
				p = "#"
			}
			fmt.Printf(p)
		}
		fmt.Printf("    \n")//, hpstr)
	}
	fmt.Printf("\n")
}

func NewUnit(t UnitType, x, y int) *Unit {
	return &Unit{
		Type: t,
		AP:   3,
		HP:   200,
		X:    x,
		Y:    y,
	}
}

func parseInput(in chan string) *Game {
	m := make([][]*Cell, 0)
	u := make([]*Unit, 0)
	w := 0
	y := 0
	for line := range in {
		row := make([]*Cell, len(line))
		if w == 0 {
			w = len(line)
		}
		for x, r := range []rune(line) {
			c := &Cell{
				Wall: r == '#',
			}
			switch r {
			case 'G':
				c.Unit = NewUnit(UnitTypeGoblin, x, y)
			case 'E':
				c.Unit = NewUnit(UnitTypeElf, x, y)
			}
			if c.Unit != nil {
				u = append(u, c.Unit)
			}
			row[x] = c
		}
		m = append(m, row)
		y++
	}
	return &Game{
		Map:   m,
		Units: u,
		mapw:  w,
		maph:  y,
	}
}

func task1(in chan string) string {
	g := parseInput(in)
	i := 0
	for ; ; i++ {
		if !g.Tick() {
			break
		}
		fmt.Printf("After %d rounds: \n", i+1)
		g.Print()
	}
	thp := 0
	for _, u := range g.Units {
		if u.HP < 0 {
			continue
		}
		thp += u.HP
	}
	fmt.Printf("Ended after %d rounds with %d total HP\n", i, thp) 
	return strconv.Itoa(i * thp)
}
