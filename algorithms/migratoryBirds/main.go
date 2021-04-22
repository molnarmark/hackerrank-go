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
	readLine(reader)

	scores := make(map[int]int)
	nums := splitInput(readLine(reader))

	for _, num := range nums {
		scores[num]++
	}

	score := 0
	bType := -999
	for i := 1; i <= 5; i++ {
		birdScore := scores[i]
		if birdScore > score && bType < i {
			score = birdScore
			bType = i
		}
	}

	fmt.Println(bType)
}
