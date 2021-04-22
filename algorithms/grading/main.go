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
	var grades []int
	data := readLine(reader)
	for data != "" {
		num, _ := strconv.Atoi(data)
		grades = append(grades, num)

		data = readLine(reader)
	}

	var finalGrades []int

	for _, grade := range grades {
		added := false
		for i := 1; i < 3; i++ {
			roundedGrade := grade + i

			if grade+3 < 40 {
				finalGrades = append(finalGrades, grade)
				added = true
				break
			}

			if roundedGrade%5 == 0 {
				finalGrades = append(finalGrades, roundedGrade)
				added = true
				break
			}
		}

		if !added {
			finalGrades = append(finalGrades, grade)
		}
	}

	for _, num := range finalGrades {
		fmt.Println(num)
	}
}
