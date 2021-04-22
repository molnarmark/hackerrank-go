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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)
	nums := splitInput(readLine(reader))

	shortest := 999
	for i, num1 := range nums {
		for j, num2 := range nums[i+1:] {
			if num1 == num2 {
				distance := (j + (i + 1)) - i
				if distance < shortest {
					shortest = distance
				}
			}
		}
	}

	if shortest == 999 {
		fmt.Println(-1)
		return
	} else {
		fmt.Println(shortest)
		return
	}
}
