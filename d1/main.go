package main

import (
	"strconv"
)

func atoiChan(in chan string) chan int {
	out := make(chan int)
	go func() {
		for line := range in {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			out <- i
		}
		close(out)
	}()
	return out
}
		

func task1(in chan string) string {
	freq := 0
	for i := range atoiChan(in) {
		freq += i
	}
	return strconv.Itoa(freq)
}

func task2(in chan string) string {
	list := make([]int, 0)
	for i := range atoiChan(in) {
		list = append(list, i)
	}
	
	freqs := make(map[int]bool)
	freq := 0
	ok := false
	for !ok {
		for _, v := range list {
			freq += v
			if _, ok = freqs[freq]; ok {
				break
			}
			freqs[freq] = true
		}
	}
	return strconv.Itoa(freq)
}