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

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)
	nums := reverseInts(splitInput(readLine(reader)))

	var strs []string

	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}

	fmt.Println(strings.Join(strs, " "))
}
