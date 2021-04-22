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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	days, _ := strconv.Atoi(readLine(reader))
	likes := 2 // day one likes are fixed
	recipients := 6

	for i := 2; i < days+1; i++ {
		newLikes := int(math.Floor(float64(recipients / 2)))
		recipients = newLikes * 3
		likes += newLikes
	}

	fmt.Println(likes)
}
