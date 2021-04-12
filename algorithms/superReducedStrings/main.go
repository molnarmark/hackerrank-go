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

func removeDups(S string) string {
	stack := []rune{}
	for _, s := range S {
		if len(stack) == 0 || len(stack) > 0 && stack[len(stack)-1] != s {
			stack = append(stack, s)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	str := readLine(reader)
	result := removeDups(str)

	if result == "" {
		result = "Empty String"
	}

	fmt.Println(result)
}
