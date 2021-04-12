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

func splitInput(input string) []int64 {
	var data []int64

	for _, val := range strings.Split(input, " ") {
		intVal, _ := strconv.ParseInt(val, 10, 64)
		data = append(data, intVal)
	}

	return data
}

var currentLine int = 0

func addLineToMatrix(matrix [][]int64, line []int64) [][]int64 {
	matrix[currentLine] = line
	currentLine++

	return matrix
}

func calcLeftToRight(matrix [][]int64, size int) int {
	column := 0
	row := 0
	sum := 0

	for i := 0; i < size; i++ {
		sum += int(matrix[column][row])
		row += 1
		column += 1
	}

	return sum
}

func calcRightToLeft(matrix [][]int64, size int) int {
	column := size - 1
	row := 0
	sum := 0

	for i := 0; i < size; i++ {
		sum += int(matrix[column][row])
		row += 1
		column -= 1
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	matrixSize, _ := strconv.ParseInt(readLine(reader), 10, 64)
	var matrix = make([][]int64, matrixSize)

	// filling up matrix
	for i := 0; i < int(matrixSize); i++ {
		line := readLine(reader)
		if line == "" {
			break
		}
		matrix = addLineToMatrix(matrix, splitInput(line))
	}

	// fmt.Println(matrix)
	// os.Exit(1)
	left := calcLeftToRight(matrix, int(matrixSize))
	right := calcRightToLeft(matrix, int(matrixSize))
	fmt.Println(math.Abs(float64(left - right)))
}
