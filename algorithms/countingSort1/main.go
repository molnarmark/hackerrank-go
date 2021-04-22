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

func countOccurrances(nums []int) []int {
	dict := make([]int, 100)

	for _, num := range nums {
		dict[num]++
	}

	return dict
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)
	nums := splitInput(readLine(reader))
	result := countOccurrances(nums)
	arr := []string{}

	for _, num := range result {
		arr = append(arr, strconv.Itoa(num))
	}

	fmt.Println(strings.Join(arr, " "))
}
