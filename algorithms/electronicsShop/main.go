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

func getHighestCombo(keyboards, drives []int, budget int) int {
	var combos []int

	for _, keyboard := range keyboards {
		for _, drive := range drives {
			if keyboard+drive <= budget {
				combos = append(combos, keyboard+drive)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(combos)))

	if len(combos) > 0 {
		return combos[0]
	} else {
		return -1
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	data := splitInput(readLine(reader))
	keyboards := splitInput(readLine(reader))
	drives := splitInput(readLine(reader))

	budget := data[0]
	sort.Sort(sort.Reverse(sort.IntSlice(keyboards)))
	sort.Sort(sort.Reverse(sort.IntSlice(drives)))

	fmt.Println(getHighestCombo(keyboards, drives, budget))

}
