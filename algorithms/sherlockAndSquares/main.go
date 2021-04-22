package main

import (
	"bufio"
	"fmt"
	"math"
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
		ranges := splitInput(readLine(reader))
		counter := 0
		i := math.Ceil(math.Sqrt(float64(ranges[0])))
		for i*i <= float64(ranges[1]) {
			counter++
			i++
		}
		fmt.Println(counter)
	}
}
