package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	aliceScore := 0
	bobScore := 0

	alice := splitInput(readLine(reader))
	bob := splitInput(readLine(reader))

	for i := 0; i < 3; i++ {
		if alice[i] > bob[i] {
			aliceScore += 1
		} else if alice[i] == bob[i] {
			//
		} else {
			bobScore += 1
		}
	}

	fmt.Println(aliceScore, bobScore)
}
