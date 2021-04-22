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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	hour, _ := strconv.Atoi(readLine(reader))
	minute, _ := strconv.Atoi(readLine(reader))

	hoursStr := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}
	minutesStr := []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
		"twenty",
		"twenty one",
		"twenty two",
		"twenty three",
		"twenty four",
		"twenty five",
		"twenty six",
		"twenty seven",
		"twenty eight",
		"twenty nine",
	}

	if minute == 0 {
		fmt.Printf("%s o' clock\n", hoursStr[hour])
		return
	}

	if minute < 15 {
		mStr := "minutes"
		if minute < 10 {
			mStr = "minute"
		}
		fmt.Printf("%s %s past %s\n", minutesStr[minute], mStr, hoursStr[hour])
		return
	}

	if minute == 15 {
		fmt.Printf("quarter past %s\n", hoursStr[hour])
		return
	}

	if minute > 15 && minute < 30 {
		fmt.Printf("%s minutes past %s\n", minutesStr[minute], hoursStr[hour])
		return
	}

	if minute == 30 {
		fmt.Printf("half past %s\n", hoursStr[hour])
		return
	}

	if minute == 45 {
		fmt.Printf("quarter to %s\n", hoursStr[hour+1])
		return
	}

	if minute > 30 && minute != 45 {
		diff := 60 - minute
		fmt.Printf("%s minutes to %s\n", minutesStr[diff], hoursStr[hour+1])
		return
	}
}
