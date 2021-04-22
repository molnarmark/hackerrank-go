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
	sentence := readLine(reader)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	message := "pangram"

	for _, letter := range alphabet {
		if !strings.Contains(sentence, string(letter)) && !strings.Contains(sentence, strings.ToUpper(string(letter))) {
			message = "not pangram"
			break
		}
	}

	fmt.Println(message)
}
