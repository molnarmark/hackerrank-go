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

	for _, word := range arr {
		dict[word]++
	}

	return dict
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)

	magazine := countOccurrances(strings.Split(readLine(reader), " "))
	target := strings.Split(readLine(reader), " ")
	answer := "Yes"

	for _, k := range target {
		magazine[k]--
		if magazine[k] == -1 {
			answer = "No"
		}
	}

	fmt.Println(answer)
}
