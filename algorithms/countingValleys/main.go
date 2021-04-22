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
	readLine(reader)
	steps := strings.Split(readLine(reader), "")

	currentLevel := 0
	previousLevel := 0
	valleysPassed := 0

	for _, step := range steps {
		if step == "U" {
			currentLevel += 1
		} else {
			currentLevel -= 1
		}

		if currentLevel == 0 && previousLevel < 0 {
			valleysPassed++
		}

		previousLevel = currentLevel
	}

	fmt.Println(valleysPassed)
}
