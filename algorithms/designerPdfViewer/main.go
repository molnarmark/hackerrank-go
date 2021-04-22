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
	abc := "abcdefghijklmnopqrstuvwxyz"
	indexes := splitInput(readLine(reader))
	word := readLine(reader)

	highest := 0
	for _, char := range strings.Split(word, "") {
		index := strings.Index(abc, char)
		abcIndex := indexes[index]
		if abcIndex > highest {
			highest = abcIndex
		}
	}

	fmt.Println(highest * len(word))
}
