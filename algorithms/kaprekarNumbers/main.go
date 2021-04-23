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

func splitAtCenterAndSum(num string) int {
	center := len(num) / 2
	num1, _ := strconv.Atoi(num[0:center])
	num2, _ := strconv.Atoi(num[center:])
	return num1 + num2
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	min, _ := strconv.Atoi(readLine(reader))
	max, _ := strconv.Atoi(readLine(reader))

	nums := []string{}
	for i := min; i <= max; i++ {
		squared := i * i

		if splitAtCenterAndSum(strconv.Itoa(squared)) == i {
			nums = append(nums, strconv.Itoa(i))
		}
	}
	if len(nums) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		fmt.Println(strings.Join(nums, " "))
	}
}
