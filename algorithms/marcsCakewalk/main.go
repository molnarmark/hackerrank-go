package main

import (
	"bufio"
	"fmt"
	"math"
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

	unsorted := splitInput(readLine(reader))
	sorted := make([]int, len(unsorted))
	copy(sorted, unsorted)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	sortedNeeded := 0
	unsortedNeeded := 0

	for i, num := range unsorted {
		unsortedNeeded += int(math.Pow(2, float64(i))) * num
	}

	for i, num := range sorted {
		sortedNeeded += int(math.Pow(2, float64(i))) * num
	}

	if sortedNeeded < unsortedNeeded {
		fmt.Println(sortedNeeded)
	} else {
		fmt.Println(unsortedNeeded)
	}
}
