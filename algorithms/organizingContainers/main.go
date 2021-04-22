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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	queryCount, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < queryCount; i++ {
		countainerCount, _ := strconv.Atoi(readLine(reader))
		containers := [][]int{}

		for j := 0; j < countainerCount; j++ {
			containers = append(containers, splitInput(readLine(reader)))
		}

		uniqueNumbers := []int{}
		containerLength := len(containers)
		for _, container := range containers {
			for _, num := range container {
				if !contains(uniqueNumbers, num) {
					uniqueNumbers = append(uniqueNumbers, num)
				}
			}

		}

		fmt.Println(uniqueNumbers, containers)
		if len(uniqueNumbers) >= containerLength {
			fmt.Println("Impossible")
		} else {
			fmt.Println("Possible")
		}
	}
}
