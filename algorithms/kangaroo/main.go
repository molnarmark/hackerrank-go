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

type Kangaroo struct {
	pos, jumpDistance int
}

func (k *Kangaroo) jump() {
	k.pos += k.jumpDistance
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	data := splitInput(readLine(reader))

	kangaroo1 := &Kangaroo{pos: data[0], jumpDistance: data[1]}
	kangaroo2 := &Kangaroo{pos: data[2], jumpDistance: data[3]}

	answer := "NO"

	for i := 0; i < 10000; i++ {
		kangaroo1.jump()
		kangaroo2.jump()

		if kangaroo1.pos == kangaroo2.pos {
			answer = "YES"
			break
		}
	}

	fmt.Println(answer)
}
