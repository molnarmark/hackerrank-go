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
	neededWord := "hackerrank"
	words, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < words; i++ {
		word := readLine(reader)
		found := ""
		p := 0
		for j := 0; j < len(word); j++ {
			if p > len(neededWord)-1 {
				break
			}
			if word[j] == neededWord[p] {
				p++
				found += string(word[j])
			}
		}

		if found == neededWord {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
