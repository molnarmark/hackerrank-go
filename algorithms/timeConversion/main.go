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
	rawTime := readLine(reader)

	splitTime := strings.Split(rawTime, ":")
	convertedHour, _ := strconv.ParseInt(splitTime[0], 10, 64)
	amOrPm := rawTime[len(rawTime)-2:]

	var hour string
	if convertedHour+12 == 24 {
		if amOrPm == "AM" {
			hour = "00"
		} else {
			hour = "12"
		}
	} else {
		if convertedHour < 12 && amOrPm == "AM" {
			hour = fmt.Sprintf("0%d", int(convertedHour))
		} else {
			hour = fmt.Sprintf("%d", int(convertedHour)+12)

		}
	}

	fmt.Printf("%s:%s\n", hour, strings.ReplaceAll(strings.Join(splitTime[1:], ":"), amOrPm, ""))
}
