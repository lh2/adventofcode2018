package main

import (
	"strconv"
	"strings"
)

type Register [4]int

type Operation struct {
	Code int
	Args [3]int
}

type Recording struct {
	Before    Register
	After     Register
	Operation Operation
}

func parseRegister(rs string) Register {
	parts := strings.Split(rs[1:len(rs)-1], ", ")
	return Register{
		mustAtoi(parts[0]),
		mustAtoi(parts[1]),
		mustAtoi(parts[2]),
		mustAtoi(parts[3]),
	}
}

func parseOperation(os string) Operation {
	parts := strings.Split(os, " ")
	return Operation{
		Code: mustAtoi(parts[0]),
		Args: [3]int{
			mustAtoi(parts[1]),
			mustAtoi(parts[2]),
			mustAtoi(parts[3]),
		},
	}
}

func parseRecordings(in chan string) []Recording {
	rs := make([]Recording, 0)
	for {
		bs := <-in
		os := <-in
		as := <-in
		if bs == "" {
			break
		}

		r := Recording{
			Before:    parseRegister(bs[8:]),
			After:     parseRegister(as[8:]),
			Operation: parseOperation(os),
		}
		rs = append(rs, r)

		<-in
	}
	return rs
}

func task1(in chan string) string {
	rs := parseRecordings(in)
	rc := 0
	for _, r := range rs {
		c := 0
		for _, op := range AllOps {
			a := r.Operation.Args
			res := op(r.Before, a[0], a[1], a[2])
			if res == r.After {
				c++
			}
			if c == 3 {
				break
			}
		}
		if c == 3 {
			rc++
		}
	}
	return strconv.Itoa(rc)
}

func filterOp(ops map[string]OpFunc, relem string) map[string]OpFunc {
	nops := make(map[string]OpFunc)
	for k, op := range ops {
		if k != relem {
			nops[k] = op
		}
	}
	return nops
}

func task2(in chan string) string {
	rs := parseRecordings(in)
	opcMap := make(map[int]map[string]OpFunc)
	for _, r := range rs {
		opc := r.Operation.Code
		if _, ok := opcMap[opc]; !ok {
			opcMap[opc] = AllOps
		}
		for opName, op := range AllOps {
			a := r.Operation.Args
			res := op(r.Before, a[0], a[1], a[2])
			if res != r.After {
				opcMap[opc] = filterOp(opcMap[opc], opName)
			}
		}

	}
	for {
		uniq := make([]string, 0)
		for _, opm := range opcMap {
			if len(opm) > 1 {
				continue
			}
			for k, _ := range opm {
				uniq = append(uniq, k)
			}
		}
		if len(uniq) == len(AllOps) {
			break
		}
		for opc, opm := range opcMap {
			if len(opm) == 1 {
				continue
			}
			for _, uop := range uniq {
				opcMap[opc] = filterOp(opcMap[opc], uop)
			}
		}
	}

	r := Register{0, 0, 0, 0}
	for line := range in {
		if line == "" {
			continue
		}
		op := parseOperation(line)
		opm := opcMap[op.Code]
		if len(opm) != 1 {
			panic("should not happen")
		}
		for _, fn := range opm {
			a := op.Args
			r = fn(r, a[0], a[1], a[2])
		}
	}
	return strconv.Itoa(r[0])
}
