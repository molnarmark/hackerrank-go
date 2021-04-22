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
	data := splitInput(readLine(reader))
	initialCost := data[0]
	reducedBy := data[1]
	minimumPrice := data[2]
	budget := data[3]

	gamesBought := 0
	previousCost := initialCost
	totalSpent := 0

	for {
		totalSpent += previousCost
		nextCost := previousCost - reducedBy
		if nextCost < minimumPrice {
			nextCost = minimumPrice
		}

		if totalSpent > budget {
			break
		}

		previousCost = nextCost
		gamesBought++
	}

	fmt.Println(gamesBought)
}
