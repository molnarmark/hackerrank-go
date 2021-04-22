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
	maxHeight := splitInput(readLine(reader))[1]
	hurdles := splitInput(readLine(reader))

	highest := 0

	for _, hurdle := range hurdles {
		if hurdle > highest {
			highest = hurdle
		}
	}

	if highest-maxHeight < 0 {
		fmt.Println(0)
		return
	}

	fmt.Println(highest - maxHeight)
}
