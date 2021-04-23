package main

import (
	"bufio"
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
	numOfCases, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < numOfCases; i++ {
		n, _ := strconv.Atoi(readLine(reader))
		a, _ := strconv.Atoi(readLine(reader))
		b, _ := strconv.Atoi(readLine(reader))
	}
}
