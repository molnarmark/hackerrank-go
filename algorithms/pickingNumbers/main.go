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

func validToAdd(subarr []int, num int) bool {
	valid := true

	if len(subarr) == 0 {
		return true
	}

	for _, arr := range subarr {
		if math.Abs(float64(arr-num)) > 1 {
			valid = false
		}
	}

	return valid
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)
	nums := splitInput(readLine(reader))
	longest := 0

	for i := 0; i < len(nums); i++ {
		subarray := []int{nums[i]}
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if validToAdd(subarray, nums[j]) {
				subarray = append(subarray, nums[j])
			}
		}
		if len(subarray) > longest {
			longest = len(subarray)
		}
	}
	fmt.Println(longest)
}
