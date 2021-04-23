package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func rotateLeft(d int, arr []string) []string {
	for ; d > 0; d-- {
		left := arr[0]
		arr = arr[1:]
		arr = append(arr, left)
	}
	return arr
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	readLine(reader)
	word := readLine(reader)
	rotation, _ := strconv.Atoi(readLine(reader))
	rotatedAlphabet := rotateLeft(rotation, alphabet)

	str := ""
	for _, letter := range word {
		r := regexp.MustCompile(`[A-Z]`)
		isUpper := r.MatchString(string(letter))
		index := indexOf(strings.ToLower(string(letter)), alphabet)
		if index == -1 {
			str += string(letter)
			continue
		}
		if isUpper {
			str += strings.ToUpper(rotatedAlphabet[index])
		} else {
			str += rotatedAlphabet[index]
		}
	}

	fmt.Println(str)
}
