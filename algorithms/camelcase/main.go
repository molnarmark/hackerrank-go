package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// utilities

func readLine(reader *bufio.Reader) string {
	text, _, _ := reader.ReadLine()
	return string(text)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	word := readLine(reader)
	regex := regexp.MustCompile(`[A-Z\s]+`)
	fmt.Println(len(regex.FindAllStringIndex(word, -1)) + 1)
}
