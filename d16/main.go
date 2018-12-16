package main

import (
	"strconv"
	"strings"
)

type Register [4]int

type Operation struct {
	OpCode int
	Args   [3]int
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
		OpCode: mustAtoi(parts[0]),
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
