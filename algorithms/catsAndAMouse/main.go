package main

import (
	"bufio"
	"fmt"
	"math"
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

func solveQuery(positions []int) string {
	catA := positions[0]
	catB := positions[1]
	mouse := positions[2]

	catADiff := math.Abs(float64(catA - mouse))
	catBDiff := math.Abs(float64(catB - mouse))

	if catADiff == catBDiff {
		return "Mouse C"
	}

	if catADiff > catBDiff {
		return "Cat B"
	} else {
		return "Cat A"
	}

	return ""
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)

	data := readLine(reader)
	var queries [][]int

	for data != "" {
		queries = append(queries, splitInput(data))
		data = readLine(reader)
	}

	for _, query := range queries {
		fmt.Println(solveQuery(query))
	}
}
