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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	var square [][]int

	for i := 0; i < 3; i++ {
		square = append(square, splitInput(readLine(reader)))
	}

	squares := [][]int{
		{8, 3, 4, 1, 5, 9, 6, 7, 2},
		{4, 3, 8, 9, 5, 1, 2, 7, 6},
		{8, 1, 6, 3, 5, 7, 4, 9, 2},
		{6, 1, 8, 7, 5, 3, 2, 9, 4},
		{2, 9, 4, 7, 5, 3, 6, 1, 8},
		{4, 9, 2, 3, 5, 7, 8, 1, 6},
		{2, 7, 6, 9, 5, 1, 4, 3, 8},
		{6, 7, 2, 1, 5, 9, 8, 3, 4},
	}

	min := float64(999)
	for i := 0; i < len(squares); i++ {
		var total float64 = 0
		for j := 0; j < len(squares[i]); j++ {
			total += math.Abs(float64(square[j/3][j%3] - squares[i][j]))
		}

		if total < min {
			min = total
		}
	}

	fmt.Println(min)
}
