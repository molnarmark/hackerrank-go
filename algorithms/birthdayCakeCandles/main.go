package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	readLine(reader)
	nums := splitInput(readLine(reader))

	sum := 0
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	tallestHeight := nums[0]

	for _, num := range nums {
		if num == tallestHeight {
			sum += 1
		} else {
			break
		}
	}

	fmt.Println(sum)
}
