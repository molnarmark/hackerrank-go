package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	date1Arr := splitInput(readLine(reader))
	date2Arr := splitInput(readLine(reader))

	returned := time.Date(date1Arr[2], time.Month(date1Arr[1]), date1Arr[0], 0, 0, 0, 0, time.UTC)
	due := time.Date(date2Arr[2], time.Month(date2Arr[1]), date2Arr[0], 0, 0, 0, 0, time.UTC)

	dayDiff := returned.Sub(due).Hours() / 24

	if dayDiff < 0 {
		fmt.Println(0)
		return
	}

	if dayDiff <= 1 && date1Arr[2] == date2Arr[2] && date1Arr[1] == date2Arr[1] {
		fmt.Println(0)
		return
	} else if dayDiff > 1 && date1Arr[2] == date2Arr[2] && date1Arr[1] == date2Arr[1] {
		fmt.Println(math.Abs(float64(15 * dayDiff)))
		return
	} else if date1Arr[2] == date2Arr[2] && date1Arr[1] != date2Arr[1] {
		fmt.Println(math.Abs(float64(500 * (date2Arr[1] - date1Arr[1]))))
		return
	} else {
		fmt.Println(10000)
		return
	}
}
