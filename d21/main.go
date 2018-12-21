package main

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Name string
	Args [3]int
}

type Register [6]int

func parseInstructions(in chan string) (is []Instruction, ipIndex int) {
	is = make([]Instruction, 0)
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
	return
}

func task1(in chan string) string {
	is, ip := parseInstructions(in)
	r := Register{}
	var idx int
	for {
		idx = r[ip]
		if idx == 28 || idx > len(is) {
			break
		}
		a := is[idx].Args
		r = AllOps[is[idx].Name](r, a[0], a[1], a[2])
		r[ip]++
	}
	return strconv.Itoa(r[2])
}
