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

func genRange(num1, num2 int) []int {
	nums := []int{}

	for i := num1; i <= num2; i++ {
		nums = append(nums, i)
	}

	return nums
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	num1, _ := strconv.Atoi(readLine(reader))
	num2, _ := strconv.Atoi(readLine(reader))

	maxXor := -999

	for _, i := range genRange(num1, num2) {
		for _, j := range genRange(num1, num2) {
			res := i ^ j
			if res > maxXor {
				maxXor = res
			}
		}
	}

	fmt.Println(maxXor)

}
