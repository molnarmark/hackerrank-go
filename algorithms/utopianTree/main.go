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

func cycle(height *int, cType int) {
	if cType == -1 {
		*height *= 2
	} else {
		*height++
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	numOfTestCases, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < numOfTestCases; i++ {
		cycles, _ := strconv.Atoi(readLine(reader))
		treeHeight := 0

		for j := 0; j <= cycles; j++ {
			if j%2 == 0 {
				cycle(&treeHeight, 1)
			} else {
				cycle(&treeHeight, -1)
			}
		}

		fmt.Println(treeHeight)
	}
}
