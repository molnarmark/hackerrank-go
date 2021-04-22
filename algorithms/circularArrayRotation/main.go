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

func rotate(nums []int, k int) {
	k = k % len(nums)
	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	data := splitInput(readLine(reader))
	_, k, q := data[0], data[1], data[2]
	arr := splitInput(readLine(reader))

	rotate(arr, k)

	for i := 0; i < q; i++ {
		num, _ := strconv.Atoi(readLine(reader))
		fmt.Println(arr[num])
	}
}
