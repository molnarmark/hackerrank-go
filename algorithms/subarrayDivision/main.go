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

func getSum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)
	nums := splitInput(readLine(reader))

	dayAndMonth := splitInput(readLine(reader))

	c := 0
	for i, _ := range nums {
		target := i + dayAndMonth[1]
		if target > len(nums) {
			break
		}
		subarr := nums[i:target]
		subarrSum := getSum(subarr)

		if subarrSum == dayAndMonth[0] {
			c++
		}
	}

	fmt.Println(c)
}
