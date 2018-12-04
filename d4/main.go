package main

import (
	"sort"
	"strconv"
	"strings"
)

func getGuardsSleepWindows(list []string) map[int][]int {
	guardsSleep := make(map[int][]int)
	guardID := 0
	start := 0
	for _, line := range list {
		msg := line[strings.Index(line, "]")+2:]
		colonPos := strings.Index(line, ":")
		minute := mustAtoi(line[colonPos+1 : colonPos+3])

		if strings.HasPrefix(msg, "Guard") {
			guardID = mustAtoi(strings.Split(msg, " ")[1][1:])
			if _, ok := guardsSleep[guardID]; !ok {
				guardsSleep[guardID] = make([]int, 60)
			}
			continue
		}
		switch msg {
		case "falls asleep":
			start = minute
		case "wakes up":
			for i := start; i < minute; i++ {
				guardsSleep[guardID][i]++
			}
		}
	}
	return guardsSleep
}

func getSleepiestGuard(guardsSleep map[int][]int) int {
	sleepiestGuard := 0
	longestSleep := 0
	for guard, times := range guardsSleep {
		time := 0
		for _, v := range times {
			time += v
		}
		if time > longestSleep {
			longestSleep = time
			sleepiestGuard = guard
		}
	}
	return sleepiestGuard
}

func getMostAsleep(sleepSchedule []int) (minute, times int) {
	for m, t := range sleepSchedule {
		if t > times {
			minute = m
			times = t
		}
	}
	return
}

func task1(in chan string) string {
	list := inAsSlice(in)
	sort.Strings(list)

	guardsSleep := getGuardsSleepWindows(list)
	sleepiestGuard := getSleepiestGuard(guardsSleep)
	sleepiestGuardSchedule := guardsSleep[sleepiestGuard]
	maxSleepMinute, _ := getMostAsleep(sleepiestGuardSchedule)

	return strconv.Itoa(sleepiestGuard * maxSleepMinute)
}

func task2(in chan string) string {
	list := inAsSlice(in)
	sort.Strings(list)

	guardsSleep := getGuardsSleepWindows(list)
	mostAsleepTimes := 0
	mostAsleepMinute := 0
	longestSleeper := 0
	for guard, schedule := range guardsSleep {
		m, t := getMostAsleep(schedule)
		if t > mostAsleepTimes {
			mostAsleepMinute = m
			mostAsleepTimes = t
			longestSleeper = guard
		}
	}

	return strconv.Itoa(longestSleeper * mostAsleepMinute)
}
