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

func splitInput(input string) []int64 {
	var data []int64

	for _, val := range strings.Split(input, " ") {
		intVal, _ := strconv.ParseInt(val, 10, 64)
		data = append(data, intVal)
	}

	return data
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	size, _ := strconv.ParseInt(readLine(reader), 10, 64)

	for i := 1; i <= int(size); i++ {
		padding := size - int64(i)
		stairs := size - padding
		fmt.Println(strings.Repeat(" ", int(padding)) + strings.Repeat("#", int(stairs)))
	}
}
