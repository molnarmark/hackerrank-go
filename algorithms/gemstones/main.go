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

func countOccurrances(word string) map[string]int {
	dict := make(map[string]int)

	for _, char := range word {
		if dict[string(char)] > 0 {
			continue
		}
		dict[string(char)] = 1
	}

	return dict
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	numOfCases, _ := strconv.Atoi(readLine(reader))
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	occurrMap := []map[string]int{}
	for i := 0; i < numOfCases; i++ {
		word := readLine(reader)
		occurrMap = append(occurrMap, countOccurrances(word))
	}
	stones := 0

	for _, char := range alphabet {
		found := 0
		for _, m := range occurrMap {
			c := string(char)
			if m[c] > 0 {
				found++
			}
		}

		if found == numOfCases {
			stones++
		}
	}
	fmt.Println(stones)
}
