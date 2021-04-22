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

func turnBook(totalPages, neededPage int, dir string) int {
	if neededPage == 1 {
		return 0
	}

	if dir == "forwards" {
		return (neededPage / 2)
	} else {
		calc := ((totalPages - neededPage) / 2)
		if calc < 0 {
			return 0
		}
		return calc
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	totalPages, _ := strconv.Atoi(readLine(reader))
	neededPage, _ := strconv.Atoi(readLine(reader))

	if neededPage == 1 {
		fmt.Println(0)
		return
	}

	if neededPage == totalPages-1 && totalPages%2 == 0 {
		fmt.Println(1)
		return
	}

	center := totalPages / 2
	if neededPage > center {
		fmt.Println(turnBook(totalPages, neededPage, "backwards"))
	} else {
		fmt.Println(turnBook(totalPages, neededPage, "forwards"))
	}

}
