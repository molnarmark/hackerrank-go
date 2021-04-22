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

func replaceAtIndex(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	gridSize, _ := strconv.Atoi(readLine(reader))

	grid := []string{}

	for i := 0; i < gridSize; i++ {
		row := readLine(reader)
		grid = append(grid, row)
	}

	for row := 1; row < gridSize-1; row++ {
		for column := 1; column < gridSize-1; column++ {
			top := string(grid[row-1][column])    // top
			center := string(grid[row][column])   // center
			bottom := string(grid[row+1][column]) // bottom
			left := string(grid[row][column-1])   // left
			right := string(grid[row][column+1])  // right

			if top < center && bottom < center && left < center && right < center {
				grid[row] = replaceAtIndex(grid[row], 'X', column)
			}
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}
}
