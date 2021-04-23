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

func rotateLeft(arr []string, d int) []string {
	for ; d > 0; d-- {
		left := arr[0]
		arr = arr[1:]
		arr = append(arr, left)
	}
	return arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	rotations := splitInput(readLine(reader))[1]
	nums := strings.Split(readLine(reader), " ")
	fmt.Println(strings.Join(rotateLeft(nums, rotations), " "))

}
