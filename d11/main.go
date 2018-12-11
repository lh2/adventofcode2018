package main

import "fmt"

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

	max := 0
	var maxx int
	var maxy int
	var maxs int
	for size := 1; size < GRID_SIZE; size++ {
		power, x, y := findMaxPower(grid, size)
		if power > max {
			max = power
			maxx = x
			maxy = y
			maxs = size
		}
	}
	return fmt.Sprintf("%d,%d,%d\n", maxx+1, maxy+1, maxs)
}
