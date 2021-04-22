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
	message := readLine(reader)
	length := len(message)
	root := math.Sqrt(float64(length))
	rows := int(math.Floor(root)) + 1
	columns := int(math.Ceil(root))

	encrypted := ""
	for i := 0; i < columns; i++ {
		for j := 0; j < rows*columns; j += columns {
			if i+j >= length {
				break
			}
			encrypted += string(message[i+j])
		}

		encrypted += " "
	}

	fmt.Println(encrypted)
}
