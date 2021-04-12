package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// utilities

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

func reverse(s []int) []int {
	a := make([]int, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func addUp4(ints []int) int {
	var sum int
	for i := 0; i < 4; i++ {
		sum += ints[i]
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	nums := splitInput(readLine(reader))
	var intNums []int

	for _, num := range nums {
		intNums = append(intNums, int(num))
	}

	sort.Ints(intNums)
	fmt.Println(addUp4(intNums), addUp4(reverse(intNums)))
}
