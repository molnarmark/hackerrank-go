package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// number of games, we dont care
	readLine(reader)

	results := splitInput(readLine(reader))
	highestScore := results[0]
	lowestScore := results[0]

	brokenHighScore := 0
	brokenLowestScore := 0

	for _, score := range results {
		if score > highestScore {
			highestScore = score
			brokenHighScore++
		}

		if score < lowestScore {
			lowestScore = score
			brokenLowestScore++
		}
	}

	fmt.Println(brokenHighScore, brokenLowestScore)
}
