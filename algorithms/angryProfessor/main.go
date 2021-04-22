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
	numOfTestCases, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < numOfTestCases; i++ {
		threshold := splitInput(readLine(reader))[1]
		arrivalTimes := splitInput(readLine(reader))
		arrivedOnTime := 0

		for _, arrivalTime := range arrivalTimes {
			if arrivalTime <= 0 {
				arrivedOnTime++
			}
		}

		if arrivedOnTime < threshold {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
