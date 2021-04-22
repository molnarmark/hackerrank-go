package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// utilities

func readLine(reader *bufio.Reader) string {
	text, _, _ := reader.ReadLine()
	return string(text)
}

func splitInput(input string) []int {
	var data []int

	for _, val := range strings.Split(input, " ") {
		intVal, _ := strconv.ParseInt(val, 10, 64)
		data = append(data, int(intVal))
	}

	return data
}

func reorder(scoreboard []int, lookingFor int) int {
	highest := math.MaxInt64
	rank := 0
	newRank := 0

	sort.Sort(sort.Reverse(sort.IntSlice(scoreboard)))
	for _, score := range scoreboard {
		if score < highest && score != highest {
			rank++
			highest = score
		}
		if score == lookingFor {
			newRank = rank
		}
	}

	return newRank
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)
	scoreboard := splitInput(readLine(reader))
	readLine(reader)
	scores := splitInput(readLine(reader))

	for _, newScore := range scores {
		scoreboard = append(scoreboard, newScore)
		jump := reorder(scoreboard, newScore)
		fmt.Println(jump)
	}
}
