package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func fib(n int) int {
	a, b := 0, 1

	for {
		n = n-1
		if n > 0 {
			a, b = b, a+b
		} else {
			break
		}


	}
	return b
}


func main() {
	//f := fibonacci()
	//printFileContent(intGen(f))
	fmt.Println(fib(50))

}