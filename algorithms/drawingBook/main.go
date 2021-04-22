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

func chunkify(totalPages int) [][]int {
	var chunks [][]int
	for i := 0; i <= totalPages-1; i = i + 2 {
		chunk := []int{i, i + 1}
		chunks = append(chunks, chunk)
	}

	return chunks
}

func solve(pageChunks [][]int, neededPage, i, direction int) int {
	page := pageChunks[i]
	// fmt.Println(page, neededPage)
	if page[0] == neededPage || page[1] == neededPage {
		if direction == -1 {
			return len(pageChunks) - i - 1
		} else {
			return i
		}
	}

	return -1
}

// TODO delete everything in here. this is terrible.

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	totalPages, _ := strconv.Atoi(readLine(reader))
	neededPage, _ := strconv.Atoi(readLine(reader))

	center := totalPages / 2

	// 1 == forwards, -1 == backwards
	direction := 1

	if neededPage > center {
		direction = -1
	}

	if neededPage == totalPages || neededPage == totalPages-1 {
		fmt.Println(0)
		return
	}

	if neededPage == totalPages-2 {
		fmt.Println(1)
		return
	}

	// fmt.Println(direction)

	pageChunks := chunkify(totalPages)
	if direction == 1 {
		for i := 0; i <= len(pageChunks); i++ {
			answer := solve(pageChunks, neededPage, i, direction)
			if answer != -1 {
				fmt.Println(answer)
				break
			}
		}
	} else {
		for i := len(pageChunks) - 1; i >= 0; i-- {
			answer := solve(pageChunks, neededPage, i, direction)
			if answer != -1 {
				fmt.Println(answer)
				break
			}
		}
	}
}
