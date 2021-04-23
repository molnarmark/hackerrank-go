package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func sortTimes(times map[int]int) []int {
	sortedKeys := []int{}

	for key := range times {
		sortedKeys = append(sortedKeys, key)
	}

	sort.Ints(sortedKeys)
	return sortedKeys
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
	numOfOrders, _ := strconv.Atoi(readLine(reader))

	deliverTimes := map[int]int{}

	for i := 0; i < numOfOrders; i++ {
		data := splitInput(readLine(reader))
		order := data[0]
		prep := data[1]
		deliverTime := order + prep
		deliverTimes[deliverTime] = i + 1
	}

	str := ""
	for _, val := range sortTimes(deliverTimes) {
		str += strconv.Itoa(deliverTimes[val]) + " "
	}

	if len(sortTimes(deliverTimes)) < numOfOrders {
		rng := genRange(1, numOfOrders)
		str = ""
		for _, val := range rng {
			str += strconv.Itoa(val) + " "
		}
		fmt.Println(str)
	} else {
		fmt.Println(str)
	}

}
