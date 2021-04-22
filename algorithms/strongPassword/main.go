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

	numbers := "0123456789"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters := "!@#$%^&*()-+"
	readLine(reader)
	password := readLine(reader)

	satisfied := 0

	for _, num := range strings.Split(numbers, "") {
		if strings.Contains(password, num) {
			satisfied++
			break
		}
	}

	for _, char := range strings.Split(lowerCase, "") {
		if strings.Contains(password, char) {
			satisfied++
			break
		}
	}

	for _, char := range strings.Split(upperCase, "") {
		if strings.Contains(password, char) {
			satisfied++
			break
		}
	}

	for _, spec := range strings.Split(specialCharacters, "") {
		if strings.Contains(password, spec) {
			satisfied++
			break
		}
	}

	fmt.Println(math.Max(float64(4-satisfied), float64(6-len(password))))
}
