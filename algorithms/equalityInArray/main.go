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

func countOccurrances(arr []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num]++
	}
	return dict
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)

	arr := splitInput(readLine(reader))
	occurranceMap := countOccurrances(arr)

	highestOccurrance := 0
	for _, m := range occurranceMap {
		if m > highestOccurrance {
			highestOccurrance = m
		}
	}

	fmt.Println(len(arr) - highestOccurrance)
}
