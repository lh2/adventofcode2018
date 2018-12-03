package main

import (
	"strconv"
)

func task1(in chan string) string {
	chksm := 0
	for line := range in {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		chksm += i
	}
	return strconv.Itoa(chksm)
}