package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	word := readLine(reader)
	n, _ := strconv.Atoi(readLine(reader))
	regex := regexp.MustCompile(`a`)

	contains := len(regex.FindAllStringIndex(word, -1))
	whole := n / len(word)
	rem := n % len(word)
	count := contains * whole

	if rem != 0 {
		contains := len(regex.FindAllStringIndex(word[0:rem], -1))
		count += contains
	}

	fmt.Println(count)
}
