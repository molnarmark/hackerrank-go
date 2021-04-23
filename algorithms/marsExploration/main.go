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
	signal := readLine(reader)

	counter := 0
	for i := 0; i < len(signal); i += 3 {
		if signal[i] != 'S' {
			counter++
		}
		if signal[i+1] != 'O' {
			counter++
		}
		if signal[i+2] != 'S' {
			counter++
		}
	}

	fmt.Println(counter)
}
