package main

import (
	"fmt"
)


func test1() {
	for i, v := range "GO语言" {
		fmt.Printf("%d %c\n", i, v)
	}
}

func test2() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for i, v := range map1 {
		fmt.Printf("%d: %s\n", i, v)
	}
}

func main() {
	test2()
}

