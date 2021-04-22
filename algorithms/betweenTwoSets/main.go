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
	a := splitInput(readLine(reader))
	b := splitInput(readLine(reader))

	c := 0
	for i := 1; i <= 100; i++ {
		factor := true

		for j := 0; j < len(a); j++ {
			if i%a[j] != 0 {
				factor = false
				break
			}
		}
		if factor {
			for j := 0; j < len(b); j++ {
				if b[j]%i != 0 {
					factor = false
					break
				}
			}
		}

		if factor {
			c++
		}

	}
	fmt.Println(c)
}
