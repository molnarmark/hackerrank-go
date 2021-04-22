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

func getMin(arr []int) int {
	lowest := 999

	for _, val := range arr {
		if val < lowest {
			lowest = val
		}
	}

	return lowest
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	testCases := splitInput(readLine(reader))[1]
	widths := splitInput(readLine(reader))

	for i := 0; i < testCases; i++ {
		data := splitInput(readLine(reader))
		min, max := data[0], data[1]+1
		fmt.Println(getMin(widths[min:max]))
	}
}
