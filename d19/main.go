package main

import (
	"strconv"
	"strings"
)

var ipIndex = 0

type Register [6]int

func (r *Register) IP() int {
	return r[ipIndex]
}

func (r *Register) SetIP(ip int) {
	r[ipIndex] = ip
}

type Instruction struct {
	Name string
	Args [3]int
}

func parseInstructions(in chan string) []Instruction {
	is := make([]Instruction, 0)
	for line := range in {
		parts := strings.Split(line, " ")
		if parts[0] == "#ip" {
			ipIndex = mustAtoi(parts[1])
			continue
		}
		is = append(is, Instruction{
			Name: parts[0],
			Args: [3]int{
				mustAtoi(parts[1]),
				mustAtoi(parts[2]),
				mustAtoi(parts[3]),
			},
		})
	}
	return is
}

func execute(is []Instruction, r Register) Register {
	var i int
	for {
		i = r.IP()
		if i > len(is) {
			break
		}
		a := is[i].Args
		r = AllOps[is[i].Name](r, a[0], a[1], a[2])
		r.SetIP(r.IP() + 1)
	}
	return r
}

func task1(in chan string) string {
	is := parseInstructions(in)
	r := Register{}
	r = execute(is, r)
	return strconv.Itoa(r[0])
}
