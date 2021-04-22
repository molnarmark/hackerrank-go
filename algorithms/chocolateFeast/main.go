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
	testCases, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < testCases; i++ {
		data := splitInput(readLine(reader))
		budget := data[0]
		cost := data[1]
		wrappersForChocolate := data[2]
		totalWrappers := 0

		eaten := 0

		for budget >= cost {
			budget -= cost
			eaten++
		}

		totalWrappers += eaten

		for totalWrappers >= -1 {
			if totalWrappers >= wrappersForChocolate {
				totalWrappers -= wrappersForChocolate
				totalWrappers++
				eaten++
			} else {
				break
			}
		}

		fmt.Println(eaten)
	}
}
