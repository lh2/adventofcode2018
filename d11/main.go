package main

import (
	"fmt"
	"sync"
)

const GRID_SIZE = 300

func buildGrid(serial int) []int {
	grid := make([]int, GRID_SIZE*GRID_SIZE)
	for x := 0; x < GRID_SIZE; x++ {
		offset := x * GRID_SIZE
		for y := 0; y < GRID_SIZE; y++ {
			rackID := (x + 1 + 10)
			powerLevel := rackID * (y + 1)
			powerLevel += serial
			powerLevel *= rackID
			powerLevel = (powerLevel - powerLevel/1000*1000) / 100
			powerLevel -= 5
			grid[offset+y] = powerLevel
		}
	}
	return grid
}

func findMaxPower(grid []int, size int) (max, maxx, maxy int) {
	for x := 0; x < GRID_SIZE-size+1; x++ {
		for y := 0; y < GRID_SIZE-size+1; y++ {
			power := 0
			for i := 0; i < size*size; i++ {
				xoffset := i / size
				yoffset := i % size
				power += grid[(x+xoffset)*GRID_SIZE+(y+yoffset)]
			}
			if power > max {
				max = power
				maxx = x
				maxy = y
			}
		}
	}
	return
}

func task1(in chan string) string {
	serial := mustAtoi(<-in)
	grid := buildGrid(serial)
	_, x, y := findMaxPower(grid, 3)
	return fmt.Sprintf("%d,%d\n", x+1, y+1)
}

func task2(in chan string) string {
	serial := mustAtoi(<-in)
	grid := buildGrid(serial)

	ch := make(chan [4]int)
	var wg sync.WaitGroup
	for size := 1; size < GRID_SIZE; size++ {
		go func(s int) {
			wg.Add(1)
			power, x, y := findMaxPower(grid, s)
			ch <- [4]int{
				power,
				x,
				y,
				s,
			}
			wg.Done()
		}(size)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	max := 0
	var maxx int
	var maxy int
	var maxs int
	for r := range ch {
		if r[0] > max {
			max = r[0]
			maxx = r[1]
			maxy = r[2]
			maxs = r[3]
		}
	}
	return fmt.Sprintf("%d,%d,%d\n", maxx+1, maxy+1, maxs)
}
