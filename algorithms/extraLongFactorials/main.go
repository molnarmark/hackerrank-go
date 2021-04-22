package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// utilities

func readLine(reader *bufio.Reader) string {
	text, _, _ := reader.ReadLine()
	return string(text)
}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	num, _ := strconv.Atoi(readLine(reader))

	fmt.Println(factorial(big.NewInt(int64(num))))
}
