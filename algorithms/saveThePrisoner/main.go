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
	numOfCases, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < numOfCases; i++ {
		data := splitInput(readLine(reader))
		numOfPrisoners := data[0]
		numOfSweets := data[1]
		startsAt := data[2]

		sol := (startsAt + numOfSweets - 1) % numOfPrisoners
		if sol == 0 {
			fmt.Println(numOfPrisoners)
		} else {
			fmt.Println(sol)
		}
	}

}
