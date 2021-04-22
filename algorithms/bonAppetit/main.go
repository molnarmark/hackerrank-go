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
	annaDidntEat := splitInput(readLine(reader))[1]
	ordered := splitInput(readLine(reader))
	charged, _ := strconv.Atoi(readLine(reader))

	sum := 0
	for i, food := range ordered {
		if i == annaDidntEat {
			continue
		}
		sum += food
	}

	perPerson := sum / 2
	if charged == perPerson {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(charged - perPerson)
	}
}
