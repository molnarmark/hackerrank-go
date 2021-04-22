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

func jump(clouds []int, position, jumps *int) bool {
	if *position+2 <= len(clouds)-1 {
		if clouds[*position+2] == 0 {
			*position += 2
			*jumps++
			return true
		}
	}

	if *position+1 <= len(clouds)-1 {
		if clouds[*position+1] == 0 {
			*position++
			*jumps++
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	readLine(reader)
	clouds := splitInput(readLine(reader))
	position := 0
	jumps := 0

	for position <= len(clouds)-1 {
		if ok := jump(clouds, &position, &jumps); !ok {
			break
		}
	}

	fmt.Println(jumps)
}
