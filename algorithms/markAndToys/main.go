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
	budget := splitInput(readLine(reader))[1]
	prices := splitInput(readLine(reader))
	sort.Ints(prices)
	toysBought := 0

	for _, price := range prices {
		if budget-price > 0 {
			toysBought++
			budget -= price
		} else {
			break
		}
	}

	fmt.Println(toysBought)
}
