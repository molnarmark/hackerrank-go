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
	atTime, _ := strconv.Atoi(readLine(reader))

	value := 3
	previousCycleStartValue := 3
	time := 1

	for time != atTime {
		value--
		if value == 0 {
			previousCycleStartValue = previousCycleStartValue * 2
			value = previousCycleStartValue
		}

		time++
	}

	fmt.Println(value)
}
