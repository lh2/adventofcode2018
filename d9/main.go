package main

import (
	"strconv"
	"strings"
)

type Marble struct {
	Value int
	Prev  *Marble
	Next  *Marble
}

func (m *Marble) Add(nm *Marble) *Marble {
	nm.Next = m.Next
	m.Next.Prev = nm
	m.Next = nm
	nm.Prev = m
	return nm
}

func (m *Marble) Remove() {
	m.Prev.Next = m.Next
	m.Next.Prev = m.Prev
	m.Prev = nil
	m.Next = nil
}

type Player struct {
	Marbles []*Marble
}

func (p *Player) Score() (score int) {
	for _, m := range p.Marbles {
		score += m.Value
	}
	return
}

func parseGameDef(def string) (pc int, mm int) {
	defp := strings.Split(def, " ")
	pc = mustAtoi(defp[0])
	mm = mustAtoi(defp[6])
	return
}

func setupGame(pc int, mm int) (bag []*Marble, players []*Player) {
	players = make([]*Player, pc)
	for i := 0; i < pc; i++ {
		players[i] = &Player{
			Marbles: make([]*Marble, 0),
		}
	}
	bag = make([]*Marble, mm+1)
	for i := 0; i <= mm; i++ {
		m := &Marble{
			Value: i,
		}
		bag[i] = m
	}
	return
}

func playGame(b []*Marble, ps []*Player) {
	cm := b[0]
	cm.Next = cm
	cm.Prev = cm
	for i, m := range b[1:] {
		cp := ps[i%len(ps)]

		if m.Value%23 > 0 {
			cm = cm.Next.Add(m)
			continue
		}

		cp.Marbles = append(cp.Marbles, m)
		rm := cm.Prev.Prev.Prev.Prev.Prev.Prev.Prev
		cm = rm.Next
		rm.Remove()
		cp.Marbles = append(cp.Marbles, rm)
	}
}

func highestScore(ps []*Player) (hs int) {
	for _, p := range ps {
		s := p.Score()
		if s > hs {
			hs = s
		}
	}
	return hs
}

func task1(in chan string) string {
	pc, mm := parseGameDef(<-in)
	b, ps := setupGame(pc, mm)
	playGame(b, ps)
	return strconv.Itoa(highestScore(ps))
}

func task2(in chan string) string {
	pc, mm := parseGameDef(<-in)
	mm *= 100
	b, ps := setupGame(pc, mm)
	playGame(b, ps)
	return strconv.Itoa(highestScore(ps))
}
