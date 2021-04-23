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

func countOccurrances(arr []string) map[string]int {
	dict := make(map[string]int)

	for _, val := range arr {
		dict[val]++
	}

	return dict
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	cases, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < cases; i++ {
		answer := "NO"
		word1 := countOccurrances(strings.Split(readLine(reader), ""))
		word2 := countOccurrances(strings.Split(readLine(reader), ""))

		for k := range word1 {
			if word2[k] != 0 {
				answer = "YES"
				break
			}
		}
		fmt.Println(answer)
	}
}
