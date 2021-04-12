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

func calculateFruitLocationFromTree(treePos int, fruits []int) []int {
	var fruitCalc []int

	for _, fruit := range fruits {
		fruitCalc = append(fruitCalc, treePos+fruit)
	}

	return fruitCalc
}

func getFruitsInRange(houseLoc []int, fruits []int) int {
	fruitsInZone := 0
	for _, fruit := range fruits {
		if fruit >= houseLoc[0] && fruit <= houseLoc[1] {
			fruitsInZone++
		}
	}

	return fruitsInZone
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	houseLoc := splitInput(readLine(reader))
	treesLoc := splitInput(readLine(reader))
	readLine(reader)
	apples := splitInput(readLine(reader))
	oranges := splitInput(readLine(reader))

	applesCalc := calculateFruitLocationFromTree(treesLoc[0], apples)
	orangesCalc := calculateFruitLocationFromTree(treesLoc[1], oranges)

	fmt.Println(getFruitsInRange(houseLoc, applesCalc))
	fmt.Println(getFruitsInRange(houseLoc, orangesCalc))
}
