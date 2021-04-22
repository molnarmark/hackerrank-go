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

func findAll(arr [][]int) int {
	highestSum := -500

	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			sum := arr[row][column] + arr[row][column+1] + arr[row][column+2] + arr[row+1][column+1] + arr[row+2][column] + arr[row+2][column+1] + arr[row+2][column+2]

			if sum > highestSum {
				highestSum = sum
			}
		}
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
