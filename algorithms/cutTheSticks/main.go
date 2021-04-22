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

func findShortest(arr []int) int {
	shortest := 100
	for _, num := range arr {
		if num < shortest {
			shortest = num
		}
	}

	return shortest
}

func subFromEach(arr []int, num int) []int {
	newArr := []int{}

	for _, origNum := range arr {
		newNum := origNum - num

		if newNum <= 0 {
			continue
		}
		newArr = append(newArr, newNum)
	}

	return newArr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)
	sticks := splitInput(readLine(reader))

	for len(sticks) > 0 {
		fmt.Println(len(sticks))
		shortest := findShortest(sticks)
		sticks = subFromEach(sticks, shortest)
	}
}
