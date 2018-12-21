package main

import "strconv"

type Point struct {
	X int
	Y int
}

type Room struct {
	Point
	Distance int
}

type Stack []*Room

func (s *Stack) Push(r *Room) {
	*s = append(*s, r)
}

func (s *Stack) Pop() *Room {
	old := *s
	if len(old) == 0 {
		return nil
	}
	*s = old[:len(old)-1]
	return old[len(old)-1]
}

func (s Stack) Peek() *Room {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}

func task1(in chan string) string {
	m := make(map[Point]*Room)
	current := &Room{}
	stack := Stack{current}
	for _, r := range []rune(<-in) {
		switch r {
		case '^':
		case '$':
			continue
		case '(':
			stack.Push(current)
		case ')':
			current = stack.Pop()
		case '|':
			current = stack.Peek()
		default:
			p := current.Point
			switch r {
			case 'N':
				p.Y--
			case 'W':
				p.X--
			case 'E':
				p.X++
			case 'S':
				p.Y++
			}
			room, ok := m[p]
			dist := current.Distance + 1
			if !ok {
				room = &Room{p, 0}
				m[p] = room
				room.Distance = dist
			} else if dist < room.Distance {
				room.Distance = dist
			}
			current = room
		}
	}
	maxd := 0
	for _, r := range m {
		if r.Distance > maxd {
			maxd = r.Distance
		}
	}
	return strconv.Itoa(maxd)
}
