package main

import "fmt"

const GRID_SIZE = 300

func task1(in chan string) string {
	serial := mustAtoi(<-in)
	grid := make([]int, GRID_SIZE*GRID_SIZE)
	for x := 0; x < GRID_SIZE; x++ {
		offset := x * GRID_SIZE
		for y := 0; y < GRID_SIZE; y++ {
			rackID := (x + 10)
			powerLevel := rackID * y
			powerLevel += serial
			powerLevel *= rackID
			powerLevel = (powerLevel - powerLevel/1000*1000) / 100
			powerLevel -= 5
			grid[offset+y] = powerLevel
		}
	}
	max := 0
	var maxx int
	var maxy int
	for x := 0; x < GRID_SIZE-2; x++ {
		for y := 0; y < GRID_SIZE-2; y++ {
			power :=
				grid[(x+0)*GRID_SIZE+(y+0)] +
				grid[(x+0)*GRID_SIZE+(y+1)] +
				grid[(x+0)*GRID_SIZE+(y+2)] +
				grid[(x+1)*GRID_SIZE+(y+0)] +
				grid[(x+1)*GRID_SIZE+(y+1)] +
				grid[(x+1)*GRID_SIZE+(y+2)] +
				grid[(x+2)*GRID_SIZE+(y+0)] +
				grid[(x+2)*GRID_SIZE+(y+1)] +
				grid[(x+2)*GRID_SIZE+(y+2)]
			if power > max {
				max = power
				maxx = x
				maxy = y
			}
		}
	}
	return fmt.Sprintf("%d,%d\n", maxx, maxy)
}
