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

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func getDiff(word []byte) string {
	diffs := ""
	for i := 0; i < len(word)-1; i++ {
		diff := int(word[i]) - int(word[i+1])
		diffs += fmt.Sprintf("%.0f", math.Abs(float64(diff)))
	}

	return diffs
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	testCases, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < testCases; i++ {
		word := readLine(reader)
		normalBytes := getDiff([]byte(word))
		reversedBytes := getDiff([]byte(reverse(word)))

		if normalBytes == reversedBytes {
			fmt.Println("Funny")
		} else {
			fmt.Println("Not Funny")
		}
	}

}
