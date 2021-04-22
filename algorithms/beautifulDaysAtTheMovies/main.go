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

func splitInput(input string) []float64 {
	var data []float64

	for _, val := range strings.Split(input, " ") {
		intVal, _ := strconv.ParseInt(val, 10, 64)
		data = append(data, float64(intVal))
	}

	return data
}

func reverseInt(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

func makeRange(min, max int) []float64 {
	a := make([]float64, max-min+1)
	for i := range a {
		a[i] = float64(min + i)
	}
	return a
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	data := splitInput(readLine(reader))
	range1 := data[0]
	range2 := data[1]
	divisor := data[2]

	beautifulDays := 0
	for _, num := range makeRange(int(range1), int(range2)) {
		reverse := reverseInt(int(num))
		div := (num - float64(reverse)) / divisor
		if float64(div) == float64(int64(div)) {
			beautifulDays++
		}
	}

	fmt.Println(beautifulDays)
}
