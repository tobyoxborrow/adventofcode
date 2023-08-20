package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(s string) []string {
	s = strings.TrimSpace(string(s))
	lines := strings.Split(s, "\n")
	sort.Strings(lines)
	return lines
}

func main() {
	lines := parseInput(input)
	guards := makeGuards(lines)

	//displayGuards(guards)

	fmt.Println("A:", SolveA(guards))
	fmt.Println("B:", SolveB(guards))
}

type Guards map[int]map[int]int

/*
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:10] falls asleep
[1518-11-01 00:50] wakes up
*/
func makeGuards(lines []string) Guards {
	guards := make(Guards)
	guardId := 0
	sleep_start := 0
	for _, line := range lines {
		minute, _ := strconv.Atoi(line[15:17])
		if line[19] == 'G' { // Guard #nn begins shift
			end := strings.Index(line, " begins")
			guardId, _ = strconv.Atoi(line[26:end])
		} else if line[19] == 'f' { // falls asleep
			sleep_start = minute
		} else { // wakes up
			sleep_end := minute
			_, ok := guards[guardId]
			if !ok {
				guards[guardId] = make(map[int]int)
			}

			for i := sleep_start; i < sleep_end; i++ {
				_, ok := guards[guardId][i]
				if !ok {
					guards[guardId][i] = 1
				} else {
					guards[guardId][i]++
				}
			}
		}

	}

	return guards
}

func displayGuards(guards Guards) {
	var sortedGuardIds []int
	for guardId := range guards {
		sortedGuardIds = append(sortedGuardIds, guardId)
	}
	sort.Ints(sortedGuardIds)
	for _, guardId := range sortedGuardIds {
		sleptMins := guards[guardId]
		var sortedMinutes []int
		for minute := range sleptMins {
			sortedMinutes = append(sortedMinutes, minute)
		}
		sort.Ints(sortedMinutes)
		fmt.Println("Guard#", guardId)
		for _, minute := range sortedMinutes {
			fmt.Printf("%02d ", minute)
		}
		fmt.Printf("\n")
		mostSleptMinute := 0
		mostSleptCount := 0
		for _, minute := range sortedMinutes {
			if sleptMins[minute] > mostSleptCount {
				mostSleptMinute = minute
				mostSleptCount = sleptMins[minute]
			}
			fmt.Printf("%02d ", sleptMins[minute])
		}
		fmt.Printf("\n")
		fmt.Println("Slept mins:", len(sleptMins))
		fmt.Printf("Most slept min: %d (%d)\n", mostSleptMinute, mostSleptCount)
	}
}

func mostSleepyGuard(guards Guards) int {
	mostSleepyGuardId := 0
	mostSleptMins := 0
	for guardId, minuteCounters := range guards {
		sleptMins := 0
		for _, m := range minuteCounters {
			sleptMins += m
		}

		if sleptMins > mostSleptMins {
			mostSleepyGuardId = guardId
			mostSleptMins = sleptMins
		}
	}
	return mostSleepyGuardId
}

func mostSleptMinute(guards Guards, guardId int) int {
	mostSleptMinute := 0
	mostSleptCount := 0
	for minute, count := range guards[guardId] {
		if count > mostSleptCount {
			mostSleptMinute = minute
			mostSleptCount = count
		}
	}
	return mostSleptMinute
}

func mostFrequentGuardMinute(guards Guards) (int, int) {
	mostFrequentGuardId := 0
	mostFrequentMinute := 0
	mostFrequentCount := 0
	for guardId, minuteCounters := range guards {
		for minute, count := range minuteCounters {
			if count > mostFrequentCount {
				mostFrequentGuardId = guardId
				mostFrequentMinute = minute
				mostFrequentCount = count
			}
		}
	}
	return mostFrequentGuardId, mostFrequentMinute
}

func SolveA(guards Guards) int {
	mostSleepyGuardId := mostSleepyGuard(guards)
	mostSleptMinute := mostSleptMinute(guards, mostSleepyGuardId)
	//fmt.Println("Guard:", mostSleepyGuardId, "Minute:", mostSleptMinute)
	return mostSleepyGuardId * mostSleptMinute
}

func SolveB(guards Guards) int {
	guardId, minute := mostFrequentGuardMinute(guards)
	//fmt.Println("Guard:", guardId, "Minute:", minute)
	return guardId * minute
}
