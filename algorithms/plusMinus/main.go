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

func splitInput(input string) []int64 {
	var data []int64

	for _, val := range strings.Split(input, " ") {
		intVal, _ := strconv.ParseInt(val, 12, 64)
		data = append(data, intVal)
	}

	return data
}

func getRatio(num int, n int64) string {
	ratio := float64(num) / float64(n)
	return fmt.Sprintf("%.5f", ratio)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	n, _ := strconv.ParseInt(readLine(reader), 10, 64)
	nums := splitInput(readLine(reader))

	positive := 0
	negative := 0
	zeros := 0

	for _, val := range nums {
		if val > 0 {
			positive++
			continue
		}
		if val < 0 {
			negative++
			continue
		}
		if val == 0 {
			zeros++
			continue
		}
	}

	fmt.Println(getRatio(positive, n))
	fmt.Println(getRatio(negative, n))
	fmt.Println(getRatio(zeros, n))
}
