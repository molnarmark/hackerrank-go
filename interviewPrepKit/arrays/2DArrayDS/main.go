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

var row, column int = 0, 0

func findHourGlass(arr [][]int) int {
	sum := 0
	var merged []int
	// fmt.Println("row: ", row, "column: ", column)
	topNums := arr[row][column : column+3]
	center := arr[row+1][column+1 : column+2]
	bottomNums := arr[row+2][column : column+3]

	merged = append(merged, topNums...)
	merged = append(merged, center...)
	merged = append(merged, bottomNums...)

	for _, num := range merged {
		sum += num
	}

	return sum
}

func findAll(arr [][]int) int {
	highestSum := -500
	for row < 5 {
		if column > 5 {
			column = 0
			row += 1
		}
		if row == 4 {
			break
		}

		glassSum := findHourGlass(arr)
		if glassSum > highestSum {
			highestSum = glassSum
		}
		column += 1
	}

	return highestSum
}

func main() {
	var arr [][]int
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	for i := 0; i < 6; i++ {
		arr = append(arr, splitInput(readLine(reader)))
	}

	fmt.Println(findAll(arr))
}
