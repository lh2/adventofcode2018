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

// source: https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
// comments removed for improved code obscurity
func primeFactors(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}
	return
}

func task2(in chan string) string {
	is := parseInstructions(in)
	r := Register{1}

	var i int
	n := 0
	samples := make([]Register, 0)
	for ; ; n++ {
		if n > 1200 {
			break
		}
		if n > 1000 {
			samples = append(samples, r)
		}

		i = r.IP()
		a := is[i].Args
		r = AllOps[is[i].Name](r, a[0], a[1], a[2])
		r.SetIP(r.IP() + 1)
	}

	changingRegisters := make(map[int]bool, 0)
	lowRegisters := make(map[int]bool, 0)
	for i, _ := range r {
		lowRegisters[i] = true
	}
	lr := samples[0]
	for _, r := range samples[1:] {
		for i, v := range r {
			if lr[i] != v {
				changingRegisters[i] = true
			}
			if v > 10000 {
				lowRegisters[i] = false
			}
		}
		lr = r
	}
	for i, v := range lowRegisters {
		if v {
			changingRegisters[i] = true
		}
	}
	ri := 15
	for i, _ := range changingRegisters {
		ri -= i
	}

	pfs := primeFactors(r[ri])
	x1 := pfs[0]
	x2 := pfs[1]
	x3 := pfs[2]

	return strconv.Itoa(1 + x1 + x2 + x3 + x1*x2 + x2*x3 + x1*x3 + x1*x2*x3)
}
